# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Копируем файлы модулей auth сервиса
COPY gateway/go.mod gateway/go.sum ./
RUN go mod download

# Копируем исходный код auth
COPY gateway/ ./

# Собираем сервис gateway
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/api-gateway ./cmd/server

# Final stage
FROM alpine:3.18

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем бинарник из builder stage
COPY --from=builder /app/api-gateway .

EXPOSE 50051

CMD ["./api-gateway"]