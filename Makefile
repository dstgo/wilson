.PHONY: install
install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install ./framework/kratosx/cmd/protoc-gen-go-errorsx
	go install ./framework/kratosx/cmd/protoc-gen-go-httpx
	go install github.com/envoyproxy/protoc-gen-validate@latest

.PHONY: lint
lint:
	golangci-lint run ./...

export MODE := debug
export OS := $(shell go env GOOS)
export ARCH := $(shell go env GOARCH)

.PHONY: build
build:
	@for service in service/*; do \
		if [ -d "$$service" ] && [ -f "$$service/Makefile" ]; then \
			echo "$$service"; \
			$(MAKE) -s -C $$service build; \
		else \
			echo "$$service does not contain Makefile, skip"; \
		fi \
	done

.PHONY: clean
clean:
	@for service in service/*; do \
		if [ -d "$$service" ] && [ -f "$$service/Makefile" ]; then \
			echo "$$service"; \
			$(MAKE) -s -C $$service clean; \
		else \
			echo "$$service does not contain Makefile, skip"; \
		fi \
	done

# proto buffer
API_DIR := ./api/proto
API_THIRD_PARTY_DIR := ./api/third_party
API_GEN_DIR := ./api/gen
API_DOC_DIR := ./api/doc
API_PB_FILES := $(shell find $(API_DIR) -name "*.proto")

.PHONY: pb
pb:
	# create dir
	mkdir -p $(API_GEN_DIR)
	# generate proto files
	protoc --proto_path=$(API_DIR)\
              --proto_path=$(API_THIRD_PARTY_DIR) \
              --go_out=$(API_GEN_DIR)\
			  --go-grpc_out=$(API_GEN_DIR)\
			  --go-httpx_out=$(API_GEN_DIR)\
			  --go-errorsx_out=$(API_GEN_DIR)\
			  --validate_out=lang=go:$(API_GEN_DIR)\
              $(API_PB_FILES)

.PHONY: doc
doc:
	# create doc dir
	mkdir -p $(API_DOC_DIR)
	# generate openapi doc
	protoc --proto_path=$(API_DIR)\
				  --proto_path=$(API_THIRD_PARTY_DIR) \
				  --openapi_out=fq_schema_naming=true,default_response=false:$(API_DOC_DIR)\
				  $(API_PB_FILES)





