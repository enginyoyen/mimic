VERSION := $(shell git describe --tags)
GITCOMMIT := $(shell git rev-parse --short HEAD)
BUILDTIME := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")


LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.GitCommit=$(GITCOMMIT) -X=main.BuildTime=$(BUILDTIME)"

build:
	@echo "  >  Building binary..."
	go build $(LDFLAGS) -o bin/mimic main.go

run:
	go run main.go
