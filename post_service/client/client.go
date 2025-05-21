package main

import (
	"context"
	"fmt"
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
	req := &post_service.ListPostsData{
		Page:      1,
		Limit:     10,
		CreatorId: "user123",
	}

	res, err := client.ListPosts(context.Background(), req)
	if err != nil {
		log.Fatalf("could not create post: %v", err)
	}
	for _, post := range res.GetPosts() {
		fmt.Println(post)
	}

}
