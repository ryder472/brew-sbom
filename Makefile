.PHONY: build test clean

build:
	go build -o bin/brew-sbom ./cmd/brew-sbom

test:
	go test ./...
	