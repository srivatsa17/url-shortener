.PHONY: build run clean test

# Clean build artifacts
clean:
	rm -rf bin/url-shortener

# Build the application
build:
	go mod download
	go build -o bin/url-shortener main.go

# Format code
fmt:
	go fmt ./...

# Run linter
lint:
	golangci-lint run ./...

# Run tests
test:
	go test -v ./...

# Run the application
run: clean build fmt lint test
	./bin/url-shortener
