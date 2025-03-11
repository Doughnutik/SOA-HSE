package main

import (
	"log"
	"user_service/database"
	"user_service/database/config"
)

func main() {
	config.LoadConfig()
	database.InitDB()
	database.RunMigrations()

	log.Println("Сервис user_service запущен")
	select {} // Блокируем выполнение, чтобы процесс не завершался
}
