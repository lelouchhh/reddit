# Используем официальный базовый образ Go
FROM golang:1.20-alpine AS build

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum, чтобы установить зависимости
COPY go.mod ./
RUN go mod download

# Копируем остальные файлы
COPY . .

# Компилируем Go-приложение
RUN go build -o app cmd/main.go

# Используем минимальный базовый образ для конечного контейнера
FROM alpine:latest

WORKDIR /root/

# Копируем скомпилированное приложение из предыдущего шага
COPY --from=build /app/app .

# Копируем файлы окружения
COPY .env .

# Определяем команду для запуска приложения
CMD ["./app"]