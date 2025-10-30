ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif


migrate-up:
	@goose postgres "user=${NOTIFICATION_POSTGRES_USER} password=${NOTIFICATION_POSTGRES_PASSWORD} dbname=${NOTIFICATION_POSTGRES_DATABASE} host=${NOTIFICATION_POSTGRES_HOST} port=${NOTIFICATION_POSTGRES_PORT} sslmode=disable" up -dir ./migrations

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
	goose postgres "user=${NOTIFICATION_POSTGRES_USER} password=${NOTIFICATION_POSTGRES_PASSWORD} dbname=${NOTIFICATION_POSTGRES_DATABASE} host=${NOTIFICATION_POSTGRES_HOST} port=${NOTIFICATION_POSTGRES_PORT} sslmode=disable" status -dir ./migrations

migration-up:
	goose postgres "user=${NOTIFICATION_POSTGRES_USER} password=${NOTIFICATION_POSTGRES_PASSWORD} dbname=${NOTIFICATION_POSTGRES_DATABASE} host=${NOTIFICATION_POSTGRES_HOST} port=${NOTIFICATION_POSTGRES_PORT} sslmode=disable" up -dir ./migrations

migration-down:
	goose postgres "user=${NOTIFICATION_POSTGRES_USER} password=${NOTIFICATION_POSTGRES_PASSWORD} dbname=${NOTIFICATION_POSTGRES_DATABASE} host=${NOTIFICATION_POSTGRES_HOST} port=${NOTIFICATION_POSTGRES_PORT} sslmode=disable" down -dir ./migrations

run:
	go run ./cmd/app/main.go


.PHONY:
	migrate \
	.up-postgres \
	.up-kafka \
	.up-service \
	up