.PHONY: gobuild gotest lint lint\:fix cl ul \
		build up in in\:localstack stop down destroy rebuild

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

# create lambda
cl:
	docker compose exec go make gobuild
	docker compose exec localstack bash scripts/create_lambda.sh
# update lambda
ul:
	docker compose exec go make gobuild
	docker compose exec localstack bash scripts/lambda/update_function.sh

# Docker
build:
	docker compose up -d --build

up:
	docker compose up -d

in:
	docker compose exec go bash

in\:localstack:
	docker compose exec localstack bash

stop:
	docker compose stop

down:
	docker compose down --remove-orphans

destroy:
	docker compose down --rmi all --volumes --remove-orphans

rebuild: destroy build