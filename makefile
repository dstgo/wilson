
.PHONY: init
init:
	go mod tidy
	go install github.com/swaggo/swag/cmd/swag@latest
	go get github.com/swaggo/gin-swagger@latest
	go get github.com/swaggo/files@latest
	go install github.com/google/wire/cmd/wire@latest
	go get github.com/google/wire/cmd/wire@latest

.PHONY: gen
gen:
	go get github.com/google/wire/cmd/wire@latest
	go generate ./...

.PHONY: build
build:
	make gen
	go build -trimpath -ldflags "-X main.Author=stranger -X main.Version=$(shell git describe --tags --always)" -o ./bin/ github.com/dstgo/wilson


.PHONY: all
all:
	make init
	make build