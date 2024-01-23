wilson_app := "github.com/dstgo/wilson/cmd/wilson"
wigfird_app := "github.com/dstgo/wilson/cmd/wigfrid"
author := github.com/dstgo
version := $(shell git describe --tags --always)
git_version := $(shell git describe --tags --always)
build_time := $(shell date +"%Y%m%d%H%M%S")
host_os := $(shell go env GOHOSTOS)
host_arch := $(shell go env GOHOSTARCH)

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
build_wilson:
	go vet ./...
	go build -trimpath \
				-ldflags="-X main.Author=$(author) -X main.Version=$(version) -X main.BuildTime=$(build_time)" \
				-o ./bin/wilson/ $(wilson_app)

build_wigfrid:
	go vet ./..
	go build -trimpath \
					-ldflags="-X main.Author=$(author) -X main.Version=$(version) -X main.BuildTime=$(build_time)" \
					-o ./bin/wilson/ $(wigfird_app)

.PHONY: gen_build
gen_build:
	make gen
	make build

.PHONY: install
install:
	make init
	make gen_build