
CLI_NAME=countdown
SERVER_NAME=server

cli:
	go build -o bin/$(CLI_NAME) cmd/cli/*.go

server:
	go build -o bin/$(SERVER_NAME) cmd/server/*.go