.PHONY: help
help:
	@echo  init	initial app dependencies
	@echo  generate	generate app template code

.PHONY: init
init:
	go mod tidy
	go install github.com/swaggo/swag/cmd/swag@latest
	go get github.com/swaggo/gin-swagger@latest
	go get github.com/swaggo/files@latest
	go install github.com/google/wire/cmd/wire@latest
	go get github.com/google/wire/cmd/wire@latest

.PHONY: generate
generate:
	go generate ./...

.PHONY: build
build:
	go build -trimpath -o ./bin/wilson.exe
