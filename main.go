package main

import (
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
)

const proxyAddress = "127.0.0.1:8080"

var backendServers = []string{
	"127.0.0.1:9091",
	"127.0.0.1:9092",
	"127.0.0.1:9093",
}

var currentServer int
var mu sync.Mutex

func main() {
	listener, err := net.Listen("tcp", proxyAddress)
	if err != nil {
		fmt.Println("Error starting proxy:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Proxy listening on", proxyAddress)

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(clientConn)
	}
}

func handleConnection(clientConn net.Conn) {
	defer clientConn.Close()

	server := handleLoadBalancer()
	fmt.Println("Forwarding request to backend:", server)

	// Read the client's request
	clientReader := make([]byte, 4096)
	n, err := clientConn.Read(clientReader)
	if err != nil {
		fmt.Println("Error reading client request:", err)
		return
	}

	// Extract method and path
	requestData := string(clientReader[:n])
	requestLines := strings.Split(requestData, "\r\n")
	if len(requestLines) < 1 {
		fmt.Println("Invalid request")
		return
	}

	// Forward request to backend
	backendConn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Println("Error connecting to backend:", err)
		clientConn.Write([]byte(fmt.Sprintf("HTTP/1.1 502 Bad Gateway\r\nContent-Type: application/json\r\n\r\n{\"error\": \"Backend %s is down\"}", server)))
		return
	}
	defer backendConn.Close()

	fmt.Println("Forwarding request to backend:\n", requestLines[0])
	backendConn.Write(clientReader[:n])

	// Read backend response
	buffer := make([]byte, 4096)
	n, err = backendConn.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Println("Error reading from backend:", err)
		return
	}

	// Send response back to client
	fmt.Println("Sending response back to client")
	clientConn.Write(buffer[:n])
}

func handleLoadBalancer() string {
	mu.Lock()
	defer mu.Unlock()

	// Round-robin selection of the backend server
	server := backendServers[currentServer]
	currentServer = (currentServer + 1) % len(backendServers)

	return server
}
