package post_service_api

import (
	"api_gateway/config"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	gen "api_gateway/post_service_api/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type PostServiceClient interface {
	CreatePost(ctx context.Context, in *gen.PostCreateData, opts ...grpc.CallOption) (*gen.PostCreateResponse, error)
	GetPost(ctx context.Context, in *gen.PostGetData, opts ...grpc.CallOption) (*gen.PostGetResponse, error)
	UpdatePost(ctx context.Context, in *gen.PostUpdateData, opts ...grpc.CallOption) (*gen.PostUpdateResponse, error)
	DeletePost(ctx context.Context, in *gen.PostDeleteData, opts ...grpc.CallOption) (*gen.PostDeleteResponse, error)
	ListPosts(ctx context.Context, in *gen.ListPostsData, opts ...grpc.CallOption) (*gen.ListPostsResponse, error)
}

var (
	ErrorIncorrectRequest = errors.New("некорректные параметры запроса")
	ErrorInternal         = errors.New("внутренняя ошибка сервиса")
	Cfg                   config.Config
)

func validateRequestBody[T any](r *http.Request, v *T) error {
	if r.Header.Get("Content-Type") != "application/json" {
		return ErrorIncorrectRequest
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("validateRequestBody\t Ошибка чтения запроса: %v", err)
		return ErrorInternal
	}
	r.Body = io.NopCloser(bytes.NewBuffer(body))
	if err := json.Unmarshal(body, v); err != nil {
		return ErrorIncorrectRequest
	}
	return nil
}

func connectToGRPCServer() (*grpc.ClientConn, error) {
	url := Cfg.PostServiceUrl
	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("connectToGRPCServer\t ошибка подключения к grpc серверу: %v", err)
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}
	return conn, nil
}

func copyGRPCResponse(w http.ResponseWriter, grpcResponse interface{}, grpcError error) {
	w.Header().Set("Content-Type", "application/json")

	responseData, err := json.Marshal(grpcResponse)
	if err != nil {
		http.Error(w, "ошибка сериализации ответа", http.StatusInternalServerError)
		log.Printf("copyGRPCResponse\t ошибка сериализации ответа: %v", err)
		return
	}

	var statusCode int = http.StatusOK
	if grpcError != nil {
		code := status.Code(grpcError)
		switch code {
		case codes.NotFound:
			statusCode = http.StatusNotFound
		case codes.InvalidArgument:
			statusCode = http.StatusBadRequest
		case codes.Internal:
			statusCode = http.StatusInternalServerError
		case codes.PermissionDenied:
			statusCode = http.StatusUnauthorized
		default:
			http.Error(w, "Неизвестная ошибка", http.StatusInternalServerError)
		}
	}
	w.WriteHeader(statusCode)
	_, err = w.Write(responseData)
	if err != nil {
		http.Error(w, "ошибка отправки ответа", http.StatusInternalServerError)
		log.Printf("copyGRPCResponse\t ошибка отправки ответа: %v", err)
	}
}
