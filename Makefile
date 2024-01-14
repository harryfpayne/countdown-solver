CLI_NAME=countdown
SERVER_NAME=server

.PHONY: all cli server

cli:
	go build -o bin/$(CLI_NAME) cmd/cli/*.go

server: frontend
	go build -o bin/$(SERVER_NAME) cmd/server/*.go