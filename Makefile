#auth-service:50051 pg 5431
#chat-service:50052    5432
#social-service:50053  5433
#users-service:50054   5434
#notification-service:50055 5435

CHAT_POSTGRES_DB := postgres-chat
CHAT_POSTGRES_USER := postgres-chat-user
CHAT_POSTGRES_PASSWORD := postgres-chat-psw

SOC_POSTGRES_DB := postgres-soc
SOC_POSTGRES_USER := postgres-soc-user
SOC_POSTGRES_PASSWORD := postgres-soc-psw

USERS_POSTGRES_DB := postgres-users
USERS_POSTGRES_USER := postgres-users-user
USERS_POSTGRES_PASSWORD := postgres-users-psw

NOTIFICATION_POSTGRES_DB := postgres-notification
NOTIFICATION_POSTGRES_USER := postgres-notification-user
NOTIFICATION_POSTGRES_PASSWORD := postgres-notification-psw

build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down

logs:
	docker-compose logs -f

clean:
	docker-compose down -v --rmi local


.migration-status-chat:
	goose postgres "user=${CHAT_POSTGRES_USER} password=${CHAT_POSTGRES_PASSWORD} dbname=${CHAT_POSTGRES_DB} host=localhost port=5432 sslmode=disable" status -dir ./chat/migrations

.migration-up-chat:
	goose postgres "user=${CHAT_POSTGRES_USER} password=${CHAT_POSTGRES_PASSWORD} dbname=${CHAT_POSTGRES_DB} host=localhost port=5432 sslmode=disable" up -dir ./chat/migrations

.migration-down-chat:
	goose postgres "user=${CHAT_POSTGRES_USER} password=${CHAT_POSTGRES_PASSWORD} dbname=${CHAT_POSTGRES_DB} host=localhost port=5432 sslmode=disable" down -dir ./chat/migrations

.migration-status-soc:
	goose postgres "user=${SOC_POSTGRES_USER} password=${SOC_POSTGRES_PASSWORD} dbname=${SOC_POSTGRES_DB} host=localhost port=5433 sslmode=disable" status -dir ./social/migrations

.migration-up-soc:
	goose postgres "user=${SOC_POSTGRES_USER} password=${SOC_POSTGRES_PASSWORD} dbname=${SOC_POSTGRES_DB} host=localhost port=5433 sslmode=disable" up -dir ./social/migrations

.migration-down-soc:
	goose postgres "user=${SOC_POSTGRES_USER} password=${SOC_POSTGRES_PASSWORD} dbname=${SOC_POSTGRES_DB} host=localhost port=5433 sslmode=disable" down -dir ./social/migrations

.migration-status-users:
	goose postgres "user=${USERS_POSTGRES_USER} password=${USERS_POSTGRES_PASSWORD} dbname=${USERS_POSTGRES_DB} host=localhost port=5434 sslmode=disable" status -dir ./users/migrations

.migration-up-users:
	goose postgres "user=${USERS_POSTGRES_USER} password=${USERS_POSTGRES_PASSWORD} dbname=${USERS_POSTGRES_DB} host=localhost port=5434 sslmode=disable" up -dir ./users/migrations

.migration-down-users:
	goose postgres "user=${USERS_POSTGRES_USER} password=${USERS_POSTGRES_PASSWORD} dbname=${USERS_POSTGRES_DB} host=localhost port=5434 sslmode=disable" down -dir ./users/migrations

.migration-status-notification:
	goose postgres "user=${NOTIFICATION_POSTGRES_USER} password=${NOTIFICATION_POSTGRES_PASSWORD} dbname=${NOTIFICATION_POSTGRES_DB} host=localhost port=5435 sslmode=disable" status -dir ./notification/migrations

.migration-up-notification:
	goose postgres "user=${NOTIFICATION_POSTGRES_USER} password=${NOTIFICATION_POSTGRES_PASSWORD} dbname=${NOTIFICATION_POSTGRES_DB} host=localhost port=5435 sslmode=disable" up -dir ./notification/migrations

.migration-down-notification:
	goose postgres "user=${NOTIFICATION_POSTGRES_USER} password=${NOTIFICATION_POSTGRES_PASSWORD} dbname=${NOTIFICATION_POSTGRES_DB} host=localhost port=5435 sslmode=disable" down -dir ./notification/migrations


migration-all-up: .migration-up-chat .migration-up-soc .migration-up-users .migration-up-notification
migration-all-down: .migration-down-chat .migration-down-soc .migration-down-users .migration-down-notification
migration-all-status: .migration-status-chat .migration-status-soc .migration-status-users .migration-status-notification

send-message-to-chat:
	go run ./chat/cmd/client/*

.PHONY: build up