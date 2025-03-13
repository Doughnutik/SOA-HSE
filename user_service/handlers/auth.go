package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
	"user_service/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUser обрабатывает регистрацию пользователя.
func RegisterUser(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.RegisterRequest

		body, err := readRequestBody(r)
		if err != nil {
			http.Error(w, "Ошибка чтения запроса", http.StatusInternalServerError)
			log.Printf("RegisterUser\t Ошибка чтения запроса: %v", err)
			return
		}

		if err := json.Unmarshal(body, &req); err != nil {
			http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
			return
		}

		// Проверяем, существует ли уже пользователь с таким логином
		var existingUserID string
		if err := db.QueryRow(context.Background(),
			"SELECT user_id FROM user_info WHERE login = $1",
			req.Login).Scan(&existingUserID); err != nil && !errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "Ошибка чтения базы данных", http.StatusInternalServerError)
			log.Printf("RegisterUser\t Ошибка проверки существования login в бд: %v", err)
			return
		} else if err == nil {
			http.Error(w, "Логин уже существует", http.StatusConflict)
			return
		}

		// Проверяем, существует ли уже пользователь с таким email
		if err := db.QueryRow(context.Background(),
			"SELECT user_id FROM user_info WHERE email = $1",
			req.Email).Scan(&existingUserID); err != nil && !errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "Ошибка чтения базы данных", http.StatusInternalServerError)
			log.Printf("RegisterUser\t Ошибка проверки существования email в бд: %v", err)
			return
		} else if err == nil {
			http.Error(w, "Email уже существует", http.StatusConflict)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Ошибка хеширования пароля", http.StatusInternalServerError)
			log.Printf("RegisterUser\t Ошибка хеширования пароля: %v", err)
			return
		}

		query := `
			INSERT INTO user_info (user_id, login, password_hash, email, registered_at, changed_at, last_login_time) 
			VALUES ($1, $2, $3, $4, $5, $5, $5)`

		newUserID := uuid.New().String() // Генерация нового UUID
		log.Printf("RegisterUser\t Создание нового пользователя %s", newUserID)

		_, err = db.Exec(context.Background(), query, newUserID, req.Login, string(hashedPassword), req.Email, time.Now())
		if err != nil {
			http.Error(w, "Ошибка сохранения пользователя", http.StatusInternalServerError)
			log.Printf("RegisterUser\t Ошибка добавления пользователя в бд: %v", err)
			return
		}

		w.WriteHeader(http.StatusOK) // Возвращаем только статус 200
	}
}

// LoginUser обрабатывает вход пользователя.
func LoginUser(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.LoginRequest

		body, err := readRequestBody(r)
		if err != nil {
			http.Error(w, "Ошибка чтения запроса", http.StatusInternalServerError)
			log.Printf("LoginUser\t Ошибка чтения запроса: %v", err)
			return
		}

		if err := json.Unmarshal(body, &req); err != nil {
			http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
			return
		}

		var storedHash string
		query := `SELECT password_hash FROM user_info WHERE login=$1`
		err = db.QueryRow(context.Background(), query, req.Login).Scan(&storedHash)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "Ошибка чтения базы данных", http.StatusInternalServerError)
			log.Printf("LoginUser\t Ошибка аутентификации: %v", err)
			return
		} else if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(req.Password)); err != nil {
			http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
			return
		}

		_, err = db.Exec(context.Background(), `UPDATE user_info SET last_login_time=$1 WHERE login=$2`, time.Now(), req.Login)
		if err != nil {
			log.Printf("Ошибка обновления last_login_time: %v", err)
		}

		//TODO нужно вернуть JWN токен
		w.WriteHeader(http.StatusOK) // Возвращаем только статус 200
	}
}
