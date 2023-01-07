.PHONY: gobuild gotest gotest\:feature gotest\:all featuretest alltest lint lint\:fix cl ul \
		build up restart in in\:localstack stop down destroy rebuild

# App
gobuild:
	mkdir -p dist
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
		go build \
			-ldflags '-w -s' \
			-o dist/greeting cmd/greeting/main.go
	zip -rj ./dist/greeting ./dist/greeting
	rm ./dist/greeting

gotest:
	go test ./app/adapter/... -count=1
	go test ./app/gateway/... -count=1
	go test ./app/infra/... -count=1

gotest\:feature:
	go test ./app/acceptance/... -count=1

gotest\:all:
	go test ./... -count=1

featuretest:
	docker compose exec localstack bash scripts/helper/upload.sh
	docker compose exec go make gotest:feature

alltest:
	docker compose exec localstack bash scripts/helper/upload.sh
	docker compose exec go make gotest:all

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
	docker compose exec localstack bash scripts/update_lambda.sh

# Docker
build:
	docker compose up -d --build

up:
	docker compose up -d

restart:
	$(MAKE) stop
	$(MAKE) up
	@sleep 15 # s3が反映されるのを待つ
	$(MAKE) cl

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