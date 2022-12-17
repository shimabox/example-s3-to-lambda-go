.PHONY: lint build up in down destroy

# App
lint:
	golangci-lint run ./... --timeout=5m

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