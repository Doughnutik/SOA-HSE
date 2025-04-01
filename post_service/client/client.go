package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	post_service "post_service/gen"
)

func main() {
	// Подключение к gRPC серверу
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Создание клиента
	client := post_service.NewPostServiceClient(conn)

	// Создание нового поста
	req := &post_service.PostCreateData{
		Title:       "My new post",
		Description: "This is a description",
		CreatorId:   "user123",
		IsPrivate:   false,
		Tags:        []string{"tag1", "tag2"},
	}

	res, err := client.CreatePost(context.Background(), req)
	if err != nil {
		log.Fatalf("could not create post: %v", err)
	}

	log.Printf("Created post with ID: %s", res.GetId())
}
