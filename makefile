
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
	go build -trimpath -ldflags "-X main.Author=$(shell git config user.name) -X main.Version=$(shell git describe --tags --always)" -o ./bin/ github.com/dstgo/wilson

.PHONY: gen_build
gen_build:
	make gen
	make build

.PHONY: install
install:
	make init
	make gen_build