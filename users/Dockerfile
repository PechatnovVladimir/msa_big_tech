# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Копируем файлы модулей chat сервиса
COPY users/go.mod users/go.sum ./
RUN go mod download

# Копируем исходный код chat
COPY users/ ./

# Собираем сервис chat
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/users-server ./cmd/app

# Final stage
FROM alpine:3.18

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем бинарник из builder stage
COPY --from=builder /app/users-server .

EXPOSE 50054

CMD ["./users-server"]
