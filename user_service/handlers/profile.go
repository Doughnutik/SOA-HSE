package handlers

import (
	"context"
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
		if errors.Is(err, DataBaseError) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
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

		if len(req.Name) == 0 && len(req.Surname) == 0 && len(req.Email) == 0 && len(req.PhoneNumber) == 0 && req.BirthDate.IsZero() {
			http.Error(w, "Нет полей для изменения", http.StatusBadRequest)
			return
		}

		var (
			newName        = req.Name
			newSurname     = req.Surname
			newBirthDate   = req.BirthDate
			newEmail       = req.Email
			newPhoneNumber = req.PhoneNumber
		)

		var (
			currentName        string
			currentSurname     string
			currentBirthDate   time.Time
			currentEmail       string
			currentPhoneNumber string
		)

		err = db.QueryRow(context.Background(),
			"SELECT name, surname, birth_date, email, phone_number FROM user_info AS t1 JOIN user_additional as t2 ON t1.user_id = t2.user_id WHERE t1.user_id=$1", userID).
			Scan(&currentName, &currentSurname, &currentBirthDate, &currentEmail, &currentPhoneNumber)

		if err != nil {
			http.Error(w, "Ошибка получения текущих данных профиля", http.StatusInternalServerError)
			log.Printf("UpdateProfile\t Ошибка получения текущих данных профиля: %v", err)
			return
		}

		if len(newName) == 0 {
			newName = currentName
		}
		if len(newSurname) == 0 {
			newSurname = currentSurname
		}
		if newBirthDate.IsZero() {
			newBirthDate = currentBirthDate
		}
		if len(newEmail) == 0 {
			newEmail = currentEmail
		}
		if len(newPhoneNumber) == 0 {
			newPhoneNumber = currentPhoneNumber
		}

		_, err = db.Exec(context.Background(),
			"UPDATE user_info SET email=$1, changed_at=$2 WHERE user_id=$3",
			newEmail, time.Now(), userID)

		if err != nil {
			http.Error(w, "Ошибка обновления профиля", http.StatusInternalServerError)
			log.Printf("UpdateProfile\t Ошибка обновления профиля в user_info: %v", err)
			return
		}

		_, err = db.Exec(context.Background(),
			"UPDATE user_additional SET name=$1, surname=$2, birth_date=$3, phone_number=$4 WHERE user_id=$5",
			newName, newSurname, newBirthDate, newPhoneNumber, userID)

		if err != nil {
			http.Error(w, "Ошибка обновления профиля", http.StatusInternalServerError)
			log.Printf("UpdateProfile\t Ошибка обновления профиля в user_additional: %v", err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
