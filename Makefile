ifneq (,$(wildcard ./.env))
    include .env
    export
endif

run:
	go run main.go

proxy:
	fly proxy 5432 -a dalle-db

migrate:
	goose -dir migrations  postgres "user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DATABASE) host=$(POSTGRES_HOST)" up

migrate-down:
	goose -dir migrations  postgres "user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DATABASE) host=$(POSTGRES_HOST)" down