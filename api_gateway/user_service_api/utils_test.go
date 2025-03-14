package user_service_api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestValidateRequestBody(t *testing.T) {
	tests := []struct {
		name          string
		requestBody   string
		contentType   string
		expectedError error
	}{
		{
			name:          "Valid UpdateProfileData JSON",
			requestBody:   `{"login": "test", "password": "12345", "birth_date": "2001-01-01", "name": "name"}`,
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
			requestBody:   `{"login": "test", "password": "12345", "email": "@gmail"}`,
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
			var requestData UpdateProfileData

			// Вызываем функцию валидации
			err := validateRequestBody(req, &requestData)

			// Проверяем результат
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("Ожидалась ошибка: %v, но получено: %v", tc.expectedError, err)
			}

			// Если JSON валидный, проверяем, правильно ли он распарсился
			if err == nil {
				var expectedData UpdateProfileData
				_ = json.Unmarshal([]byte(tc.requestBody), &expectedData)
				if requestData != expectedData {
					t.Errorf("Ожидалась структура %+v, но получено %+v", expectedData, requestData)
				}
			}
		})
	}
}
