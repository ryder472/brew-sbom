.PHONY: build test lint clean

build:
	go build -o bin/brew-sbom ./cmd/brew-sbom

test:
	go test ./...

lint:
	golangci-lint run

clean:
	rm -rf bin/
