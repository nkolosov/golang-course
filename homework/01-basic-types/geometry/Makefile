.PHONY: test lint build clean

test:
	go test ./...

lint:
	golangci-lint run ./...

build:
	go build -o bin/geometry-app ./cmd

clean:
	rm -rf bin/

install-deps:
	go mod tidy

run: build
	./bin/geometry-app