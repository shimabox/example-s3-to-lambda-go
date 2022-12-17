.PHONY: build up in down destroy

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