APP_NAME := sURL
BIN_DIR := build

EXT :=
ifeq ($(OS),Windows_NT)
	EXT := .exe
endif

CONFIG_PKG := $(shell go list ./configs)

PKG := ./cmd/unified
BIN := $(BIN_DIR)/$(APP_NAME)$(EXT)

BOT_BIN := $(BIN_DIR)/$(APP_NAME)-bot$(EXT)
BOT_PKG := ./cmd/bot
BOT_LDFLAGS := -X '$(CONFIG_PKG).botRequired=true'

BACKEND_BIN := $(BIN_DIR)/$(APP_NAME)-backend$(EXT)
BACKEND_PKG := ./cmd/backend
BACKEND_LDFLAGS := -X '$(CONFIG_PKG).backendRequired=true'

.PHONY: all build build-bot build-backend run run-bot run-backend fmt vet deps

all: build

build:
	@mkdir -p $(BIN_DIR)
	@go build -ldflags="$(BOT_LDFLAGS) $(BACKEND_LDFLAGS)" -o $(BIN) $(PKG)

build-bot:
	@mkdir -p $(BIN_DIR)
	@go build -ldflags="$(BOT_LDFLAGS)" -o $(BOT_BIN) $(BOT_PKG)

build-backend:
	@mkdir -p $(BIN_DIR)
	@go build -ldflags="$(BACKEND_LDFLAGS)" -o $(BACKEND_BIN) $(BACKEND_PKG)

run:
	@go run -ldflags="$(BOT_LDFLAGS) $(BACKEND_LDFLAGS)" $(PKG)

run-bot:
	@go run -ldflags="$(BOT_LDFLAGS)" $(BOT_PKG)

run-backend:
	@go run -ldflags="$(BACKEND_LDFLAGS)" $(BACKEND_PKG)

clean:
	@rm -rf $(BIN_DIR)

deps:
	@go mod tidy

fmt:
	@go fmt $(PKG)

vet:
	@go vet $(PKG)
