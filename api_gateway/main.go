package main

import (
	"api_gateway/config"
	"api_gateway/server"
	"api_gateway/user_service_api"
	"log"
	"net/http"
	"os"
)

func main() {
	logFile, err := os.OpenFile("api_gateway.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
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
	user_service_api.Cfg = *cfg

	srv := server.NewServer(cfg)
	log.Printf("Сервер запущен на порту %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, srv.Router))
}
