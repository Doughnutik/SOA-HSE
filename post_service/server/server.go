package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	post_service "post_service/gen"
)

// Сервисы, которые реализуют интерфейс из post_service_grpc.pb.go
type server struct {
	post_service.UnimplementedPostServiceServer
}

// Реализация метода CreatePost
func (s *server) CreatePost(ctx context.Context, req *post_service.PostCreateData) (*post_service.PostCreateResponse, error) {
	// Логика создания поста
	log.Println("Creating post:", req.GetTitle())

	// Вернуть ответ с ID
	return &post_service.PostCreateResponse{
		Id: "generated_id", // Пример ID
	}, nil
}

func main() {
	// Создание TCP listener
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Создание gRPC сервера
	grpcServer := grpc.NewServer()

	// Регистрация сервера
	post_service.RegisterPostServiceServer(grpcServer, &server{})

	// Включение reflection для gRPC серверов
	reflection.Register(grpcServer)

	// Запуск сервера
	log.Println("Starting server on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
