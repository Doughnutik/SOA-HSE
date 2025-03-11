# Используем официальный образ Go для сборки
FROM golang:1.23.7 AS builder

# Установим рабочую директорию
WORKDIR /app

# Копируем все исходные файлы в контейнер
COPY . .

# Скачиваем зависимости
RUN go mod tidy

# Собираем приложение
RUN go build -o user_service .

# Создаем второй этап сборки для уменьшения размера контейнера
FROM alpine:latest

# Копируем бинарник из предыдущего этапа
COPY --from=builder /app/user_service /usr/local/bin/user_service

# Указываем, что контейнер будет слушать на порту 8081
EXPOSE 8081

# Команда для запуска приложения
CMD ["/usr/local/bin/user_service"]