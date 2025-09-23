# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Копируем файлы модулей chat сервиса
COPY social/go.mod social/go.sum ./
RUN go mod download

# Копируем исходный код chat
COPY social/ ./

# Собираем сервис chat
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/social-server ./cmd/app

# Final stage
FROM alpine:3.18

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем бинарник из builder stage
COPY --from=builder /app/social-server .

EXPOSE 50053

CMD ["./social-server"]
