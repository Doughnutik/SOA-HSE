package post_service_api

import (
	gen "api_gateway/post_service_api/gen"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DummyResponse struct {
	Message string `json:"message"`
}

func TestCopyGRPCResponse_Success(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := DummyResponse{Message: "ok"}

	copyGRPCResponse(rr, resp, nil)

	if rr.Code != http.StatusOK {
		t.Errorf("ожидался статус %d, получен %d", http.StatusOK, rr.Code)
	}

	var result DummyResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &result); err != nil {
		t.Fatalf("не удалось распарсить JSON: %v", err)
	}

	if result.Message != "ok" {
		t.Errorf("ожидалось сообщение 'ok', получено '%s'", result.Message)
	}
}

func TestCopyGRPCResponse_NotFound(t *testing.T) {
	rr := httptest.NewRecorder()
	err := status.Error(codes.NotFound, "not found")

	copyGRPCResponse(rr, nil, err)

	if rr.Code != http.StatusNotFound {
		t.Errorf("ожидался статус %d, получен %d", http.StatusNotFound, rr.Code)
	}
}

func TestCopyGRPCResponse_InternalError(t *testing.T) {
	rr := httptest.NewRecorder()
	err := status.Error(codes.Internal, "internal error")

	copyGRPCResponse(rr, nil, err)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("ожидался статус %d, получен %d", http.StatusInternalServerError, rr.Code)
	}
}

func TestCopyGRPCResponse_InvalidJSON(t *testing.T) {
	rr := httptest.NewRecorder()

	// Каналы не сериализуются в JSON — вызовет ошибку
	invalid := make(chan int)

	copyGRPCResponse(rr, invalid, nil)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("ожидался статус %d, получен %d", http.StatusInternalServerError, rr.Code)
	}
}

func TestCopyGRPCResponse_UnknownError(t *testing.T) {
	rr := httptest.NewRecorder()
	err := errors.New("some other error") // не gRPC

	copyGRPCResponse(rr, nil, err)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("ожидался статус %d, получен %d", http.StatusInternalServerError, rr.Code)
	}
}

func TestValidateRequestBody(t *testing.T) {
	tests := []struct {
		name          string
		requestBody   string
		contentType   string
		expectedError error
	}{
		{
			name:          "Valid PostCreateData JSON",
			requestBody:   `{"login": "test", "password": "12345", "title": "title", "description": "description", "tags": ["tag1", "tag2"], "isPrivate": true}`,
			contentType:   "application/json",
			expectedError: nil,
		},
		{
			name:          "Invalid UpdateProfileData JSON",
			requestBody:   `{"login": "test", "pass": "1234", "phone_number": "80123123"`, // Незакрытая скобка
			contentType:   "application/json",
			expectedError: ErrorIncorrectRequest,
		},
		{
			name:          "Empty body",
			requestBody:   ``,
			contentType:   "application/json",
			expectedError: ErrorIncorrectRequest,
		},
		{
			name:          "Incorrect Content-Type",
			requestBody:   `{"login": "test", "password": "12345", "title": "title", "description": "description", "tags": ["tag1", "tag2"], "isPrivate": true}`,
			contentType:   "text/plain",
			expectedError: ErrorIncorrectRequest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Создаем тестовый HTTP-запрос
			req := &http.Request{
				Header: make(http.Header),
				Body:   io.NopCloser(strings.NewReader(tc.requestBody)),
			}
			req.Header.Set("Content-Type", tc.contentType)

			// Структура для хранения распарсенных данных
			var requestData gen.PostCreateData

			// Вызываем функцию валидации
			err := validateRequestBody(req, &requestData)

			// Проверяем результат
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("Ожидалась ошибка: %v, но получено: %v", tc.expectedError, err)
			}
		})
	}
}
