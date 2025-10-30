# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Копируем файлы модулей chat сервиса
COPY notification/go.mod notification/go.sum ./
RUN go mod download

# Копируем исходный код chat
COPY notification/ ./

# Собираем сервис chat
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/notification-server ./cmd/app

# Final stage
FROM alpine:3.18

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем бинарник из builder stage
COPY --from=builder /app/notification-server .

EXPOSE 50055

CMD ["./notification-server"]
