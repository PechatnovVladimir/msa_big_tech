include vendor.proto.mk

SERVICE_NAME := chat

POSTGRES_DB := postgres-$(SERVICE_NAME)
POSTGRES_USER := postgres-$(SERVICE_NAME)-user
POSTGRES_PASSWORD := postgres-$(SERVICE_NAME)-psw


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
	go list -m -f '{{.Version}}' github.com/PechatnovVladimir/msa_big_tech/$(SERVICE_NAME)@HW2

#generate-protoc: .protoc-generate
generate: .buf-generate

# Форматирование protobuf файлов
proto-format:
	$(info run buf format...)
	buf format ./proto -w

vercommit: .ver-commit

#
run:
	CurrentUser="5d4490c9-e106-4d25-a186-3a59d6d7334e" go run ./cmd/app/main.go

pg-up:
	docker-compose up -d

pg-down:
	docker-compose down

migration-status:
	goose postgres "user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} host=localhost port=5432 sslmode=disable" status -dir ./migrations

migration-up:
	goose postgres "user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} host=localhost port=5432 sslmode=disable" up -dir ./migrations

migration-down:
	goose postgres "user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} host=localhost port=5432 sslmode=disable" down -dir ./migrations


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
