all: build
.PHONY: all build

VERSION := $(shell git rev-parse --short HEAD)
BUILDTIME := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

GOLDFLAGS += -X 'main.BuildVersion=$(VERSION)'
GOLDFLAGS += -X 'main.BuildTime=$(BUILDTIME)'

build:
	@CGO_ENABLED=0 go build -ldflags="-s -w $(GOLDFLAGS)" -o ./bin/arrowhead ./cmd/main.go

install:
	cp ./bin/arrowhead /usr/local/bin

test:
	@cd pkg/parsers; grc go test -v --race
	@cd pkg/rpc; grc go test -v --race
	@cd pkg/security/keytool; grc go test -v --race
	@cd pkg/security/openssl; grc go test -v --race
