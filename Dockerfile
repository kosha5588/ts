# Используем официальный образ Go на базе более новой версии Debian
FROM golang:1.20-buster AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY ./backend/ ./backend/

# Компилируем приложение
WORKDIR /app/backend/cmd
RUN go build -o my-web-app main.go

# Используем чистый Debian для запуска
FROM debian:bookworm-slim

# Устанавливаем необходимые пакеты
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Копируем скомпилированное приложение из предыдущего этапа
COPY --from=builder /app/backend/cmd/my-web-app /usr/local/bin/my-web-app

# Указываем переменную окружения для конфигурации
ENV DATABASE_URL=postgres://user:password@db:5432/mydb?sslmode=disable

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["my-web-app"]
