CONTEXT?=dev:review:test

.PHONY: serve local unlocal commit frontend frontend-build translation bookingserviceVuku rabbitmqInDocker

build:
	go build main.go

serve:
	DEBUG=1 CONTEXT=$(CONTEXT) go run main.go serve

update:
	go get flamingo.me/flamingo/v3
