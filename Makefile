ifneq ("$(wildcard .env)", "")
	include .env
endif
export

.PHONY: *

test:
	go test -cover -race  ./...

lint:
	golangci-lint run

build:
	go build -o banner_rotator

run:
	docker-compose up -d
