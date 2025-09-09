# Dockerfile
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Копируем go модули
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Устанавливаем swag и генерируем документацию
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g cmd/server/main.go

# Собираем бинарный файл
RUN go build -o main ./cmd/server/main.go

# Финальный stage
FROM alpine:latest

# Устанавливаем зависимости
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем бинарный файл и необходимые файлы
COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./main"]