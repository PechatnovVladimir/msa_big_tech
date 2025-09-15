# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Копируем файлы модулей auth сервиса
COPY auth/go.mod auth/go.sum ./
RUN go mod download

# Копируем исходный код auth
COPY auth/ ./

# Собираем сервис auth
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/auth-server ./cmd/server

# Final stage
FROM alpine:3.18

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем бинарник из builder stage
COPY --from=builder /app/auth-server .

EXPOSE 50051

CMD ["./auth-server"]