# basic
MODULE := $(shell go list ./cmd/...)
APP_NAME := service-$(notdir $(MODULE))
GIT_VERSION := $(shell \
  if git rev-parse --is-inside-work-tree > /dev/null 2>&1 && [ -n "$$(git log -1 --format=%h 2>/dev/null)" ]; then \
    git tag --sort=-version:refname | sed -n 1p; \
  else \
    echo "v0.0.0"; \
  fi)
BUILD_TIME := $(shell date +"%Y%m%d%H%M%S")-$(shell git log -1 --format=%h 2>/dev/null || echo "no-commit")


# build
MODE := debug
BIN_DIR := $(shell pwd)/bin
HOST_OS := $(shell go env GOHOSTOS)
HOST_ARCH := $(shell go env GOHOSTARCH)
OS := $(if $(OS),$(OS),$(HOST_OS))
ARCH := $(if $(ARCH),$(ARCH),$(HOST_ARCH))
LD_FLAGS := $(nullstring)

ifeq ($(OS),)
  OS := $(HOST_OS)
endif

ifeq ($(ARCH),)
  ARCH := $(HOST_ARCH)
endif

# reduce binary size at release mode
ifeq ($(MODE), release)
	LD_FLAGS += -s -w
endif

# inject meta info
ifneq ($(APP_NAME), $(nullstring))
	LD_FLAGS += -X main.Name=$(APP_NAME)
endif
ifneq ($(BUILD_TIME), $(nullstring))
	LD_FLAGS += -X main.BuildTime=$(BUILD_TIME)
endif
ifneq ($(GIT_VERSION), $(nullstring))
	LD_FLAGS += -X main.Version=$(GIT_VERSION)
endif

# binary extension
EXT = $(nullstring)
ifeq ($(OS), windows)
	EXT = .exe
endif


.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: build
build:
	# go lint
	golangci-lint run ./...

	# prepare target environment $(os)/$(arch)
	go env -w GOOS=$(if $(OS),$(OS),$(HOST_OS))
	go env -w GOARCH=$(if $(ARCH),$(ARCH),$(HOST_ARCH))

	# build go module
	go build -trimpath \
		-ldflags="$(LD_FLAGS)" \
		-o $(BIN_DIR)/$(MODE)/$(APP_NAME)-$(OS)-$(ARCH)/$(APP_NAME)$(EXT) \
		$(MODULE)

	echo "$(MODULE)@$(GIT_VERSION)"

	# resume host environment $(host_os)/$(host_arch)
	go env -w GOOS=$(HOST_OS)
	go env -w GOARCH=$(HOST_ARCH)

.PHONY:
clean:
	rm -rf $(BIN_DIR)
