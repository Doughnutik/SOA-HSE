package post_service_api

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	gen "api_gateway/post_service_api/gen"

	"github.com/go-chi/chi/v5"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := connectToGRPCServer()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	var req gen.PostCreateData
	err = validateRequestBody(r, &req)
	if errors.Is(err, ErrorIncorrectRequest) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if errors.Is(err, ErrorInternal) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := gen.NewPostServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	grpcResponse, grpcError := client.CreatePost(ctx, &req)
	if grpcError != nil {
		http.Error(w, "ошибка запроса к gRPC серверу", http.StatusInternalServerError)
		log.Printf("createPostHandler\t ошибка запроса к gRPC серверу: %v", grpcError)
		return
	}
	copyGRPCResponse(w, grpcResponse, grpcError)
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := connectToGRPCServer()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	var req gen.PostGetData
	err = validateRequestBody(r, &req)
	if errors.Is(err, ErrorIncorrectRequest) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if errors.Is(err, ErrorInternal) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	postID := chi.URLParam(r, "post_id")
	req.Id = postID
	client := gen.NewPostServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	grpcResponse, grpcError := client.GetPost(ctx, &req)
	if grpcError != nil {
		http.Error(w, "ошибка запроса к gRPC серверу", http.StatusInternalServerError)
		log.Printf("getPostHandler\t ошибка запроса к gRPC серверу: %v", grpcError)
		return
	}
	copyGRPCResponse(w, grpcResponse, grpcError)
}

func UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := connectToGRPCServer()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	var req gen.PostUpdateData
	err = validateRequestBody(r, &req)
	if errors.Is(err, ErrorIncorrectRequest) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if errors.Is(err, ErrorInternal) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	postID := chi.URLParam(r, "post_id")
	req.Id = postID
	client := gen.NewPostServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	grpcResponse, grpcError := client.UpdatePost(ctx, &req)
	if grpcError != nil {
		http.Error(w, "ошибка запроса к gRPC серверу", http.StatusInternalServerError)
		log.Printf("updatePostHandler\t ошибка запроса к gRPC серверу: %v", grpcError)
		return
	}
	copyGRPCResponse(w, grpcResponse, grpcError)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := connectToGRPCServer()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	var req gen.PostDeleteData
	err = validateRequestBody(r, &req)
	if errors.Is(err, ErrorIncorrectRequest) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if errors.Is(err, ErrorInternal) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	postID := chi.URLParam(r, "post_id")
	req.Id = postID
	client := gen.NewPostServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	grpcResponse, grpcError := client.DeletePost(ctx, &req)
	if grpcError != nil {
		http.Error(w, "ошибка запроса к gRPC серверу", http.StatusInternalServerError)
		log.Printf("deletePostHandler\t ошибка запроса к gRPC серверу: %v", grpcError)
		return
	}
	copyGRPCResponse(w, grpcResponse, grpcError)
}

func ListPostsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := connectToGRPCServer()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	var req gen.ListPostsData
	err = validateRequestBody(r, &req)
	if errors.Is(err, ErrorIncorrectRequest) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if errors.Is(err, ErrorInternal) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := gen.NewPostServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	grpcResponse, grpcError := client.ListPosts(ctx, &req)
	if grpcError != nil {
		http.Error(w, "ошибка запроса к gRPC серверу", http.StatusInternalServerError)
		log.Printf("listPostsHandler\t ошибка запроса к gRPC серверу: %v", grpcError)
		return
	}
	copyGRPCResponse(w, grpcResponse, grpcError)
}
