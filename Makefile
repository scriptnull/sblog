.PHONY: build install test

test:
	go test -v ./...

build:
	go build cmd/sblog.go

install: build
	mv sblog ~/s/
