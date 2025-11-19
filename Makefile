include ./.env.example
export


.PHONY: build run docker up init-db

build:
	go build -o bin/sd-svc-gateway ./cmd/gateway

run:
	go run ./cmd/gateway

test:
	go test ./tests/... -v

docker:
	docker build -t sd-svc-gateway:local .
