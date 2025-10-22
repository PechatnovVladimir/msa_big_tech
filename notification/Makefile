ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif


migrate-up:
	@goose postgres "user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} host=127.0.0.1 port=5435 sslmode=disable" up -dir ./migrations

.up-postgres:
	@docker-compose up postgresql -d

.up-kafka:
	@docker-compose up kafka kafka-ui -d

.up-service:
	@docker-compose up service --build -d

up: .up-postgres .up-kafka

down:
	@docker-compose down

#Cоздать миграцию goose create <migration_name> sql -dir=./migrations
migration-status:
	goose postgres "user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} host=localhost port=5435 sslmode=disable" status -dir ./migrations

migration-up:
	goose postgres "user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} host=localhost port=5435 sslmode=disable" up -dir ./migrations

migration-down:
	goose postgres "user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} host=localhost port=5435 sslmode=disable" down -dir ./migrations

run:
	go run ./cmd/app/main.go


.PHONY:
	migrate \
	.up-postgres \
	.up-kafka \
	.up-service \
	up