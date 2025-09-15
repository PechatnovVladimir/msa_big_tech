# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Копируем файлы модулей chat сервиса
COPY chat/go.mod chat/go.sum ./
RUN go mod download

# Копируем исходный код chat
COPY chat/ ./

# Собираем сервис chat
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/chat-server ./cmd/server

# Final stage
FROM alpine:3.18

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем бинарник из builder stage
COPY --from=builder /app/chat-server .

EXPOSE 50052

CMD ["./chat-server"]
