CLI_NAME=countdown
SERVER_NAME=server

NODE_PACKAGE_MANAGER=yarn

.PHONY: all deps cli server frontend

deps:
	cd client; $(NODE_PACKAGE_MANAGER) install

cli:
	go build -o bin/$(CLI_NAME) cmd/cli/*.go

server: frontend
	go build -o bin/$(SERVER_NAME) cmd/server/*.go

frontend:
	cd client; $(NODE_PACKAGE_MANAGER) start