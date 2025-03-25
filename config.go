package main

import "time"

type Config struct {
	Servers []ServerConfig `yaml:"servers"`
}

type ServerConfig struct {
	Address string        `yaml:"address"`
	Timeout TimeoutConfig `yaml:"timeout"`
}

type TimeoutConfig struct {
	Connect time.Duration `yaml:"connect"`
	Read    time.Duration `yaml:"read"`
	Write   time.Duration `yaml:"write"`
}
