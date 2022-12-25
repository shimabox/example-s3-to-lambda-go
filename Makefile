.PHONY: gobuild gotest lint lint\:fix build up in down destroy

# App
gobuild:
	mkdir -p dist
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
		go build \
			-ldflags '-w -s' \
			-o dist/greeting_handler cmd/app/greeting/main.go
	zip -rj ./dist/greeting_handler ./dist/greeting_handler
	rm ./dist/greeting_handler

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