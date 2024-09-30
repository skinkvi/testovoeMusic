#!/bin/sh

# Ждем, пока база данных станет доступной
until pg_isready -h db -p 5432 -U postgres; do
  echo "Waiting for database to be ready..."
  sleep 2
done

# Выполняем миграции
goose -dir ./db/migrations up

# Запускаем основное приложение
exec go run ./cmd/music/main.go
