package handlers

import (
	"errors"
	"log"
	"net/http"
	"time"
	"user_service/models"
	"user_service/queries"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUser обрабатывает регистрацию пользователя.
func RegisterUser(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := ParseRequestBody[models.RegistrationData](w, r)
		if err != nil {
			return
		}
		// Проверяем, существует ли уже пользователь с таким логином
		ok, err := queries.IsLoginExists(db, req.Login)
		if err != nil {
			http.Error(w, "Ошибка чтения базы данных", http.StatusInternalServerError)
			log.Printf("RegisterUser\t Ошибка проверки существования login в бд: %v", err)
			return
		} else if ok {
			http.Error(w, "Логин уже существует", http.StatusConflict)
			return
		}

		// Проверяем, существует ли уже пользователь с таким email
		ok, err = queries.IsEmailExists(db, req.Email)
		if err != nil {
			http.Error(w, "Ошибка чтения базы данных", http.StatusInternalServerError)
			log.Printf("RegisterUser\t Ошибка проверки существования email в бд: %v", err)
			return
		} else if ok {
			http.Error(w, "Email уже существует", http.StatusConflict)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Ошибка хеширования пароля", http.StatusInternalServerError)
			log.Printf("RegisterUser\t Ошибка хеширования пароля: %v", err)
			return
		}

		err = queries.RegisterUser(db, req, string(hashedPassword))
		if err != nil {
			http.Error(w, "Ошибка сохранения пользователя", http.StatusInternalServerError)
			log.Printf("RegisterUser\t Ошибка добавления пользователя: %v", err)
			return
		}

		w.WriteHeader(http.StatusOK) // Возвращаем только статус 200
	}
}

// LoginUser обрабатывает вход пользователя.
func LoginUser(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := ParseRequestBody[models.AuthData](w, r)
		if err != nil {
			return
		}

		err = AuthenticateUser(db, *req)
		if errors.Is(err, DataBaseError) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		err = queries.UpdateLastLoginTime(db, req.Login, time.Now())
		if err != nil {
			log.Printf("Ошибка обновления last_login_time: %v", err)
		}

		w.WriteHeader(http.StatusOK)
	}
}
