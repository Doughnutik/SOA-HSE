package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
	"user_service/models"
	"user_service/queries"

	"github.com/jackc/pgx/v5/pgxpool"
)

// GetProfile возвращает профиль пользователя
func GetProfile(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := ParseRequestBody[models.AuthData](w, r)
		if err != nil {
			return
		}

		err = AuthenticateUser(db, *req)
		if errors.Is(err, ErrorDataBase) {
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

		// Получаем профиль пользователя
		profile, err := queries.GetPersonData(db, req.Login)
		if err != nil {
			http.Error(w, "Ошибка получения профиля", http.StatusInternalServerError)
			log.Printf("GetProfile\t Ошибка получения профиля из бд: %v", err)
			return
		}

		response, err := json.Marshal(profile)
		if err != nil {
			http.Error(w, "Ошибка сериализации данных", http.StatusInternalServerError)
			log.Printf("GetProfile\t Ошибка сериализации данных: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}

// UpdateProfile обновляет профиль пользователя
func UpdateProfile(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := ParseRequestBody[models.UpdateProfileData](w, r)
		if err != nil {
			return
		}

		err = AuthenticateUser(db, models.AuthData{Login: req.Login, Password: req.Password})
		if errors.Is(err, ErrorDataBase) {
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

		err = queries.UpdatePersonData(db, *req)
		if err != nil {
			switch err {
			case queries.ErrorDataBase:
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Printf("UpdateProfile\t %v", err)
			case queries.ErrorEmptyRequest:
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
