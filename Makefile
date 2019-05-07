GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BIN_FOLDER=bin/
APP_NAME=mockingbirb
BINARY_UNIX=$(APP_NAME)_unix
CONTEXT?=dev
VERSION_NUMBER?=latest

.PHONY: serve

all: test build
build:
	cd src && $(GOBUILD) -o ../$(BIN_FOLDER)$(APP_NAME) -v
build-linux:
	cd src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ../$(BIN_FOLDER)$(BINARY_UNIX) -v
test:
	$(GOTEST) ./... -v
clean:
	cd src && $(GOCLEAN)
	rm -f $(BIN_FOLDER)$(APP_NAME)
	rm -f $(BIN_FOLDER)$(BINARY_UNIX)
serve:
	cd src && DEBUG=1 CONTEXT=$(CONTEXT) go run main.go serve
update-flamingo:
	go get flamingo.me/flamingo/v3
container: build-linux containerize
containerize:
	docker build -t $(APP_NAME):$(VERSION_NUMBER) .
mockingbirb-dev: Dockerfile
	make build
	docker build --force-rm=true -t $(APP_NAME)-dev:$(VERSION_NUMBER) -f Dockerfile .
