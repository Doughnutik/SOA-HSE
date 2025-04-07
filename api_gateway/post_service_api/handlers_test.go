package post_service_api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePostHandler(t *testing.T) {
	Cfg.PostServiceUrl = "127.0.0.1:8082"
	tests := []struct {
		name               string
		requestBody        string
		contentType        string
		expectedStatusCode int
	}{
		{
			name:               "Valid Create Request",
			requestBody:        `{"login": "login", "password": "password", "title": "Test Post", "description": "Some content", "creatorId": "123", "isPrivate": true, "tags": ["go", "test"]}`,
			contentType:        "application/json",
			expectedStatusCode: 200,
		},
		{
			name:               "Invalid Create Request",
			requestBody:        `{"title": "Broken Post", "creator_id": }`,
			contentType:        "application/json",
			expectedStatusCode: 400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "http://127.0.0.1:8081/posts", bytes.NewBufferString(tt.requestBody))
			if err != nil {
				t.Fatalf("Ошибка при создании запроса: %v", err)
			}
			req.Header.Set("Content-Type", tt.contentType)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(CreatePostHandler)
			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatusCode {
				t.Errorf("Ожидался статус код %v, но получили %v", tt.expectedStatusCode, rr.Code)
			}
		})
	}
}

func TestListPostsHandler(t *testing.T) {
	Cfg.PostServiceUrl = "127.0.0.1:8082"
	tests := []struct {
		name               string
		requestBody        string
		contentType        string
		expectedStatusCode int
	}{
		{
			name:               "Valid List Request",
			requestBody:        `{"login": "login", "password": "password"}`,
			contentType:        "application/json",
			expectedStatusCode: 200,
		},
		{
			name:               "Invalid Create Request",
			requestBody:        `{"title": "Broken Post", "creator_id": }`,
			contentType:        "application/json",
			expectedStatusCode: 400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "http://127.0.0.1:8081/posts/list", bytes.NewBufferString(tt.requestBody))
			if err != nil {
				t.Fatalf("Ошибка при создании запроса: %v", err)
			}
			req.Header.Set("Content-Type", tt.contentType)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(CreatePostHandler)
			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatusCode {
				t.Errorf("Ожидался статус код %v, но получили %v", tt.expectedStatusCode, rr.Code)
			}
		})
	}
}
