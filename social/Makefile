ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

include vendor.proto.mk

SERVICE_NAME := social

# Путь до завендореных protobuf файлов
VENDOR_PROTO_PATH := $(CURDIR)/vendor.protobuf

# Путь до protobuf файлов
PROTO_PATH := $(CURDIR)/proto/api/$(SERVICE_NAME)/v1

# Путь до сгенеренных .pb.go файлов
PKG_PROTO_PATH := $(CURDIR)/pkg

.bin-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/yoheimuta/protolint/cmd/protolint@latest

# генерация .go файлов с помощью protoc (пока оставляю на всякий случай, основно способ через buf ниже)
.protoc-generate:
	mkdir -p $(PKG_PROTO_PATH)
	protoc -I $(VENDOR_PROTO_PATH) --proto_path=$(CURDIR) \
	--go_out=$(PKG_PROTO_PATH) --go_opt paths=source_relative \
	--go-grpc_out=$(PKG_PROTO_PATH) --go-grpc_opt paths=source_relative \
	--grpc-gateway_out=$(PKG_PROTO_PATH) --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true \
	$(PROTO_PATH)/service.proto \
	$(PROTO_PATH)/messages.proto

	protoc -I $(VENDOR_PROTO_PATH) --proto_path=$(CURDIR) \
	--openapiv2_out=../gateway/swagger --openapiv2_opt logtostderr=true \
	$(PROTO_PATH)/service.proto \
	$(PROTO_PATH)/messages.proto

# генерация .go файлов с помощью buf
.buf-generate:
	buf generate

.ver-commit:
	go list -m -f '{{.Version}}' github.com/PechatnovVladimir/msa_big_tech/$(SERVICE_NAME)@HW4-dev

#generate-protoc: .protoc-generate
generate: .buf-generate

# Форматирование protobuf файлов
proto-format:
	$(info run buf format...)
	buf format ./proto -w

vercommit: .ver-commit

run:
	CurrentUser="9def0b15-7ee9-425c-a9b3-ba875dfc6d11"  go run ./cmd/app/main.go

pg-up:
	docker-compose up -d

pg-down:
	docker-compose down

#Cоздать миграцию goose create <migration_name> sql -dir=./migrations
migration-status:
	goose postgres "user=${SOC_POSTGRES_USER} password=${SOC_POSTGRES_PASSWORD} dbname=${SOC_POSTGRES_DATABASE} host=${SOC_POSTGRES_HOST} port=${SOC_POSTGRES_PORT} sslmode=disable" status -dir ./migrations

migration-up:
	goose postgres "user=${SOC_POSTGRES_USER} password=${SOC_POSTGRES_PASSWORD} dbname=${SOC_POSTGRES_DATABASE} host=${SOC_POSTGRES_HOST} port=${SOC_POSTGRES_PORT} sslmode=disable" up -dir ./migrations

migration-down:
	goose postgres "user=${SOC_POSTGRES_USER} password=${SOC_POSTGRES_PASSWORD} dbname=${SOC_POSTGRES_DATABASE} host=${SOC_POSTGRES_HOST} port=${SOC_POSTGRES_PORT} sslmode=disable" down -dir ./migrations


.PHONY: \
	.bin-deps \
	.protoc-generate \
	.tidy \
	.proto-lint \
	proto-format \
	vendor \
	vendor.protobuf \
	generate \
	build \
	lint
