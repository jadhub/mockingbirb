CONTEXT?=main
GOBIN=$(shell pwd)/bin
GOFILES=$(wildcard *.go)
APPNAME=mockingbirb
VERSION_NUMBER?=latest

.PHONY: serve

build:
	go build $(APPNAME).go
	@echo "Building $(GOFILES) to ./bin for $(APPNAME)"
    @GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o bin/$(APPNAME) $(GOFILES)

serve:
	DEBUG=1 CONTEXT=$(CONTEXT) go run $(APPNAME).go serve

update:
	go get flamingo.me/flamingo/v3

mockingbirb-dev: Dockerfile
	make build
	docker build --force-rm=true -t $(APPNAME)-dev:$(VERSION_NUMBER) -f Dockerfile .