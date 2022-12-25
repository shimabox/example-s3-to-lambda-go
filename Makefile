.PHONY: gotest lint lint\:fix build up in down destroy

# App
gotest:
	go test ./...

lint:
	golangci-lint run ./... --timeout=5m
lint\:fix:
	golangci-lint run ./... --fix

# Docker
build:
	docker compose up -d --build

up:
	docker compose up -d

in:
	docker compose exec go bash

down:
	docker compose down --remove-orphans

destroy:
	docker compose down --rmi all --volumes --remove-orphans