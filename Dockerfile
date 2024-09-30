FROM golang:1.22.5-alpine

WORKDIR /app

# Устанавливаем goose
RUN apk add --no-cache git && \
    go install github.com/pressly/goose/v3/cmd/goose@latest

# Устанавливаем утилиты PostgreSQL
RUN apk add --no-cache postgresql-client

# Добавляем goose в PATH
ENV PATH="/go/bin:${PATH}"

COPY . .

RUN go mod download

# Устанавливаем переменные окружения для goose
ENV GOOSE_DRIVER=postgres
ENV GOOSE_DBSTRING="postgres://postgres:postgres@db:5432/music_db?sslmode=disable"

# Копируем скрипт запуска
COPY entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

ENTRYPOINT ["/app/entrypoint.sh"]

