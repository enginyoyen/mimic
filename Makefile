VERSION := $(shell git describe --tags)
GITCOMMIT := $(shell git rev-parse --short HEAD)
BUILDTIME := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

LDFLAGS=-ldflags "-X=github.com/enginyoyen/mimic/pkg/cmd.Version=$(VERSION) -X=github.com/enginyoyen/mimic/pkg/cmd.GitCommit=$(GITCOMMIT) -X=github.com/enginyoyen/mimic/pkg/cmd.BuildTime=$(BUILDTIME)"

build:
	@echo "  >  Building binary..."
	go build $(LDFLAGS) -o bin/mimic main.go

run:
	go run main.go
