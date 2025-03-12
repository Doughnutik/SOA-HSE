package main

import (
	"log"
	"net/http"
	"os"
	"user_service/config"
	"user_service/database"
	"user_service/server"
)

func main() {
	logFile, err := os.OpenFile("user_service.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Ошибка открытия log файла: %v", err)
	}
	log.SetOutput(logFile)
	log.Print("\n\n")

	// Загружаем конфигурацию
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Ошибка загрузки конфигурации")
	}

	// Подключаемся к базе данных
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Ошибка подключения к БД")
	}
	defer db.Close()

	// Создание таблиц в БД
	err = database.RunMigrations(db)
	if err != nil {
		log.Fatal("Ошибка создания таблиц в БД")
	}

	// Запускаем сервер
	srv := server.NewServer(cfg, db)
	log.Printf("Сервер запущен на порту %s", cfg.Port)
	log.Fatal(http.ListenAndServe(cfg.Host+":"+cfg.Port, srv.Router))
}
