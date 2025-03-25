# Contributing to Go Load Balancing Reverse Proxy

👍 First off, thanks for taking the time to contribute!

## Code of Conduct

This project and everyone participating in it is governed by our [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check the issue list as you might find out that you don't need to create one. When you are creating a bug report, please include as many details as possible:

* Use a clear and descriptive title
* Describe the exact steps which reproduce the problem
* Provide specific examples to demonstrate the steps
* Describe the behavior you observed after following the steps
* Explain which behavior you expected to see instead and why
* Include logs if relevant

### Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues. Create an issue and provide the following information:

* Use a clear and descriptive title
* Provide a step-by-step description of the suggested enhancement
* Provide specific examples to demonstrate the steps
* Describe the current behavior and explain why it should be enhanced
* Explain why this enhancement would be useful

### Pull Requests

1. Fork the repo and create your branch from `main`
2. If you've added code that should be tested, add tests
3. If you've changed APIs, update the documentation
4. Ensure the test suite passes
5. Make sure your code follows the existing style
6. Write a [good commit message](https://chris.beams.io/posts/git-commit/)

## Development Process

1. Clone the repository
```bash
git clone https://github.com/Ameriq8/reverse-proxy.git
```

2. Create a branch
```bash
git checkout -b feature/amazing-feature
```

3. Set up development environment
```bash
go mod download
```

4. Make your changes and test them
```bash
go test ./...
```

5. Commit your changes
```bash
git commit -m 'Add some amazing feature'
```

## Style Guidelines

### Git Commit Messages

* Use the present tense ("Add feature" not "Added feature")
* Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
* Limit the first line to 72 characters or less
* Reference issues and pull requests liberally after the first line

### Go Style Guide

* Follow the official [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
* Use `gofmt` to format your code
* Write meaningful comments
* Keep functions small and focused
* Use meaningful variable names

## Additional Notes

### Issue and Pull Request Labels

* `bug` - Something isn't working
* `enhancement` - New feature or request
* `documentation` - Improvements or additions to documentation
* `help wanted` - Extra attention is needed
* `good first issue` - Good for newcomers

## Recognition

Contributors who have their PRs merged will be added to the README.md contributors section.

Thank you for contributing to Go Load Balancing Reverse Proxy! 🚀
