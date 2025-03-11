# Используем официальный образ PostgreSQL
FROM postgres:latest

# Копируем файл инициализации (например, SQL скрипты для создания таблиц) в контейнер
# Этот шаг необязателен, если ваши миграции будут выполняться отдельно
COPY schema.sql /docker-entrypoint-initdb.d/

# Указываем переменные окружения для конфигурации базы данных
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=YaGnI
ENV POSTGRES_DB=user_service

# При необходимости можно указать дополнительные порты
EXPOSE 5432