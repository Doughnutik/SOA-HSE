package main

import (
	"log"
	"net"
	"os"
	"post_service/config"
	"post_service/database"
	"post_service/gen"
	"post_service/server"

	"google.golang.org/grpc"
)

func main() {
	logFile, err := os.OpenFile("post_service.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
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
	listener, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	serv := grpc.NewServer()
	gen.RegisterPostServiceServer(serv, server.NewPostServiceServer())

	log.Printf("gRPC server is running on port %s", cfg.Port)
	if err := serv.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
