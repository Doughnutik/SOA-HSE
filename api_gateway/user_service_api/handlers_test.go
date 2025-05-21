package user_service_api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterHandler(t *testing.T) {
	Cfg.UserServiceUrl = "http://127.0.0.1:8080"
	tests := []struct {
		name               string
		requestBody        string
		contentType        string
		expectedStatusCode int
	}{
		{
			name:               "Valid Registation Request",
			requestBody:        `{"login": "cotttt", "password": "dog", "email": "cottttdog@mail.com"}`,
			contentType:        "application/json",
			expectedStatusCode: 200,
		},
		{
			name:               "Invalid Registation Request",
			requestBody:        `{"log": "cat", "pass": "dog"`,
			contentType:        "application/json",
			expectedStatusCode: 400,
		},
		{
			name:               "AlreadyExists Registation Request",
			requestBody:        `{"login": "catt", "password": "dog", "email": "cattdog@mail.com"}`,
			contentType:        "application/json",
			expectedStatusCode: 409,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "http://127.0.0.1:8081/register", bytes.NewBufferString(tt.requestBody))
			if err != nil {
				t.Fatalf("Ошибка при создании запроса: %v", err)
			}
			req.Header.Set("Content-Type", tt.contentType)

			// Создаем новый записывающий ответ для хендлера
			rr := httptest.NewRecorder()

			// Вызываем хендлер
			handler := http.HandlerFunc(RegisterHandler)
			handler.ServeHTTP(rr, req)

			// Проверяем статус код
			if rr.Code != tt.expectedStatusCode {
				t.Errorf("Ожидался статус код %v, но получили %v", tt.expectedStatusCode, rr.Code)
			}
		})
	}

}

func TestLoginHandler(t *testing.T) {
	Cfg.UserServiceUrl = "http://127.0.0.1:8080"
	tests := []struct {
		name               string
		requestBody        string
		contentType        string
		expectedStatusCode int
	}{
		{
			name:               "Valid Login Request",
			requestBody:        `{"login": "cat", "password": "dog"}`,
			contentType:        "application/json",
			expectedStatusCode: 200,
		},
		{
			name:               "Invalid Registation Request",
			requestBody:        `{"log": "cat", "pass": "dog"`,
			contentType:        "application/json",
			expectedStatusCode: 400,
		},
		{
			name:               "WrongLogin Registation Request",
			requestBody:        `{"login": "cotik", "password": "dogggy"}`,
			contentType:        "application/json",
			expectedStatusCode: 401,
		},
		{
			name:               "WrongPassword Registation Request",
			requestBody:        `{"login": "cat", "password": "doggy"}`,
			contentType:        "application/json",
			expectedStatusCode: 401,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "http://127.0.0.1:8081/login", bytes.NewBufferString(tt.requestBody))
			if err != nil {
				t.Fatalf("Ошибка при создании запроса: %v", err)
			}
			req.Header.Set("Content-Type", tt.contentType)

			// Создаем новый записывающий ответ для хендлера
			rr := httptest.NewRecorder()

			// Вызываем хендлер
			handler := http.HandlerFunc(LoginHandler)
			handler.ServeHTTP(rr, req)

			// Проверяем статус код
			if rr.Code != tt.expectedStatusCode {
				t.Errorf("Ожидался статус код %v, но получили %v", tt.expectedStatusCode, rr.Code)
			}
		})
	}

}

func TestUpdateProfile(t *testing.T) {
	Cfg.UserServiceUrl = "http://127.0.0.1:8080"
	tests := []struct {
		name               string
		requestBody        string
		contentType        string
		expectedStatusCode int
	}{
		{
			name: "Valid Update Request",
			requestBody: `{"login": "cat", "password": "dog", "email": "catdog@mail.com",
			"name": "cat", "surname": "dog", "phone_number": "89999999999", "birth_date": "2000-01-01"}`,
			contentType:        "application/json",
			expectedStatusCode: 200,
		},
		{
			name:               "NotAllData Update Request",
			requestBody:        `{"login": "cat", "password": "dog", "phone_number": "8000000000", "birth_date": "2001-02-02"}`,
			contentType:        "application/json",
			expectedStatusCode: 200,
		},
		{
			name:               "Invalid Update Request",
			requestBody:        `{"login": "cat", "password" "dog", "phonumber": "8000000000", "birth_date": "2002-02-02"}`,
			contentType:        "application/json",
			expectedStatusCode: 400,
		},
		{
			name:               "NotExists Update Request",
			requestBody:        `{"login": "cottttttt", "password": "dog"}`,
			contentType:        "application/json",
			expectedStatusCode: 401,
		},
		{
			name:               "WrongAuthData Update Request",
			requestBody:        `{"login": "cat", "password": "doggggyyyy", "email": "cattdog@mail.com"}`,
			contentType:        "application/json",
			expectedStatusCode: 401,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("PUT", "http://127.0.0.1:8081/profile", bytes.NewBufferString(tt.requestBody))
			if err != nil {
				t.Fatalf("Ошибка при создании запроса: %v", err)
			}
			req.Header.Set("Content-Type", tt.contentType)

			// Создаем новый записывающий ответ для хендлера
			rr := httptest.NewRecorder()

			// Вызываем хендлер
			handler := http.HandlerFunc(UpdateProfile)
			handler.ServeHTTP(rr, req)

			// Проверяем статус код
			if rr.Code != tt.expectedStatusCode {
				t.Errorf("Ожидался статус код %v, но получили %v", tt.expectedStatusCode, rr.Code)
			}
		})
	}

}

func TestGetProfile(t *testing.T) {
	Cfg.UserServiceUrl = "http://127.0.0.1:8080"
	tests := []struct {
		name               string
		requestBody        string
		contentType        string
		expectedStatusCode int
	}{
		{
			name:               "Valid Get Request",
			requestBody:        `{"login": "cat", "password": "dog"}`,
			contentType:        "application/json",
			expectedStatusCode: 200,
		},
		{
			name:               "Invalid Update Request",
			requestBody:        `{"login": "cat", "password" "dog"}`,
			contentType:        "application/json",
			expectedStatusCode: 400,
		},
		{
			name:               "NotExists Update Request",
			requestBody:        `{"login": "cottttttt", "password": "dog"}`,
			contentType:        "application/json",
			expectedStatusCode: 401,
		},
		{
			name:               "WrongAuthData Update Request",
			requestBody:        `{"login": "cat", "password": "doggggyyyy", "email": "cattdog@mail.com"}`,
			contentType:        "application/json",
			expectedStatusCode: 401,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "http://127.0.0.1:8081/profile", bytes.NewBufferString(tt.requestBody))
			if err != nil {
				t.Fatalf("Ошибка при создании запроса: %v", err)
			}
			req.Header.Set("Content-Type", tt.contentType)

			// Создаем новый записывающий ответ для хендлера
			rr := httptest.NewRecorder()

			// Вызываем хендлер
			handler := http.HandlerFunc(GetProfile)
			handler.ServeHTTP(rr, req)

			// Проверяем статус код
			if rr.Code != tt.expectedStatusCode {
				t.Errorf("Ожидался статус код %v, но получили %v", tt.expectedStatusCode, rr.Code)
			}
		})
	}

}

// curl -X 'PUT' \
//   'http://127.0.0.1:8081/posts/20250407201522' \
//   -H 'accept: */*' \
//   -H 'Content-Type: application/json' \
//   -d '{
//   "login": "name",
//   "password": "password",
//   "title": "new_title",
//   "description": "new_description",
//   "tags": ["first", "second"],
//   "isPrivate": 0
// }'