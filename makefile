wilson_app := "github.com/dstgo/wilson/cmd/wilson"
wigfrid_app := "github.com/dstgo/wilson/cmd/wigfrid"
author := github.com/dstgo
version := $(shell git describe --tags --always)
git_version := $(shell git describe --tags --always)
build_time := $(shell date +"%Y%m%d%H%M%S")
host_os := $(shell go env GOHOSTOS)
host_arch := $(shell go env GOHOSTARCH)
# target protobuf files will be generated at api directory
target_proto_files := $(shell find ./internal/service/proto/api/ -name *.proto)

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

.PHONY: gen_build
gen_build:
	make gen
	make build

gen_pb:
	protoc --proto_path=./internal/service/proto/api/ \
		   --proto_path=./internal/service/proto/third_party/ \
		   --go_out=paths=source_relative:./internal/service \
		   --go-grpc_out=paths=source_relative:./internal/service \
		   $(target_proto_files)

.PHONY: install
install:
	make init
	make gen_build

.PHONY: build_wilson
build_wilson:
	go vet ./...
	go build -trimpath \
				-ldflags="-X main.Author=$(author) -X main.Version=$(version) -X main.BuildTime=$(build_time)" \
				-o ./bin/wilson/ $(wilson_app)

.PHONY: build_wigfrid
build_wigfrid:
	go vet ./...
	go build -trimpath \
					-ldflags="-X main.Author=$(author) -X main.Version=$(version) -X main.BuildTime=$(build_time)" \
					-o ./bin/wigfrid/ $(wigfrid_app)