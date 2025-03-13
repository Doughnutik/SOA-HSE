package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"user_service/models"
	"user_service/queries"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrorDataBase = errors.New("ошибка базы данных")
)

// readRequestBody читает тело запроса и возвращает его в виде массива байтов.
func readRequestBody(r *http.Request) ([]byte, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return body, nil
}

// writeJSONResponse отправляет JSON-ответ клиенту.
func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Ошибка сериализации ответа", http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func ParseRequestBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	var req T

	body, err := readRequestBody(r)
	if err != nil {
		http.Error(w, "Ошибка чтения запроса", http.StatusInternalServerError)
		log.Printf("ParseRequestBody\t Ошибка чтения запроса: %v", err)
		return nil, err
	}

	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
		return nil, err
	}
	return &req, nil
}

func AuthenticateUser(db *pgxpool.Pool, authData models.AuthData) error {
	storedHash, err := queries.TakePasswordByLogin(db, authData.Login)
	if err != nil {
		log.Printf("AuthenticateUser\t Ошибка аутентификации: %v", err)
		return ErrorDataBase
	}

	if len(storedHash) == 0 || bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(authData.Password)) != nil {
		return fmt.Errorf("неверный логин или пароль")
	}

	return nil // Аутентификация успешна
}
