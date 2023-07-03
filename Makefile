BINARY_NAME=gobank

build:
	@go build -o bin/$(BINARY_NAME) 

run: build
	@./bin/$(BINARY_NAME)

test:
	@go test -v ./...