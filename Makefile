APP_NAME := sURL
BIN_DIR := build

MAJOR := 1
MINOR := 0
PATCH := 0
VERSION := v$(MAJOR).$(MINOR).$(PATCH)

UPDATED_AT := $(shell date +%y%m%d)
BRANCH ?= $(shell echo $${GITHUB_REF_NAME:-$$(git rev-parse --abbrev-ref HEAD 2>/dev/null || echo "local")})
COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "000000")

EXT :=
ifeq ($(OS),Windows_NT)
	EXT := .exe
endif

CONFIG_PKG := $(shell go list ./configs)

LDFLAGS := \
-X 'main.Version=$(VERSION)' \
-X 'main.Branch=$(BRANCH)' \
-X 'main.Commit=$(COMMIT)' \
-X 'main.UpdatedAt=$(UPDATED_AT)'

PKG := ./cmd/unified
BIN := $(BIN_DIR)/$(APP_NAME)-$(VERSION)$(EXT)

BOT_BIN := $(BIN_DIR)/$(APP_NAME)-bot-$(VERSION)$(EXT)
BOT_PKG := ./cmd/bot
BOT_LDFLAGS := -X '$(CONFIG_PKG).botRequired=true'

BACKEND_BIN := $(BIN_DIR)/$(APP_NAME)-backend-$(VERSION)$(EXT)
BACKEND_PKG := ./cmd/backend
BACKEND_LDFLAGS := -X '$(CONFIG_PKG).backendRequired=true'

.PHONY: all build build-bot build-backend run run-bot run-backend fmt vet deps

all: build

build:
	@mkdir -p $(BIN_DIR)
	@go build -ldflags="$(LDFLAGS) $(BOT_LDFLAGS) $(BACKEND_LDFLAGS)" -o $(BIN) $(PKG)

build-bot:
	@mkdir -p $(BIN_DIR)
	@go build -ldflags="$(LDFLAGS) $(BOT_LDFLAGS)" -o $(BOT_BIN) $(BOT_PKG)

build-backend:
	@mkdir -p $(BIN_DIR)
	@go build -ldflags="$(LDFLAGS) $(BACKEND_LDFLAGS)" -o $(BACKEND_BIN) $(BACKEND_PKG)

run:
	@go run -ldflags="$(LDFLAGS) $(BOT_LDFLAGS) $(BACKEND_LDFLAGS)" $(PKG)

run-bot:
	@go run -ldflags="$(LDFLAGS) $(BOT_LDFLAGS)" $(BOT_PKG)

run-backend:
	@go run -ldflags="$(LDFLAGS) $(BACKEND_LDFLAGS)" $(BACKEND_PKG)

clean:
	@rm -rf $(BIN_DIR)

deps:
	@go mod tidy

fmt:
	@go fmt $(PKG)

vet:
	@go vet $(PKG)
