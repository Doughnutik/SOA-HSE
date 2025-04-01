package queries

import (
	"context"
	"errors"
	"log"
	"time"
	"user_service/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrorEmptyRequest = errors.New("пустой запрос")
	ErrorDataBase     = errors.New("ошибка базы данных")
)

func RegisterUser(db *pgxpool.Pool, req *models.RegistrationData, hashedPassword string) error {
	queryUserInfo := `
		INSERT INTO user_info (user_id, login, password_hash, registered_at, changed_at, last_login_time) 
		VALUES ($1, $2, $3, $4, $4, $4)`
	queryUserAdditional := `
		INSERT INTO user_additional (user_id, email) VALUES ($1, $2)`

	newUserID := uuid.New().String() // Генерация нового UUID
	log.Printf("registerUser\t Создание нового пользователя %s", newUserID)

	_, err := db.Exec(context.Background(), queryUserInfo, newUserID, req.Login, hashedPassword, time.Now())
	if err != nil {
		return err
	}
	_, err = db.Exec(context.Background(), queryUserAdditional, newUserID, req.Email)
	if err != nil {
		return err
	}
	return nil
}

func IsLoginExists(db *pgxpool.Pool, login string) (bool, error) {
	var existingUserID string
	err := db.QueryRow(context.Background(), "SELECT user_id FROM user_info WHERE login = $1", login).Scan(&existingUserID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return false, err
	} else if err == nil {
		return true, nil
	}
	return false, nil
}

func IsEmailExists(db *pgxpool.Pool, email string) (bool, error) {
	var existingUserID string
	err := db.QueryRow(context.Background(), "SELECT user_id FROM user_additional WHERE email = $1", email).Scan(&existingUserID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return false, err
	} else if err == nil {
		return true, nil
	}
	return false, nil
}

func TakePasswordByLogin(db *pgxpool.Pool, login string) (string, error) {
	var storedHash string
	query := `SELECT password_hash FROM user_info WHERE login=$1`
	err := db.QueryRow(context.Background(), query, login).Scan(&storedHash)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return "", err
	} else if errors.Is(err, pgx.ErrNoRows) {
		return "", nil
	}
	return storedHash, nil
}

func UpdateLastLoginTime(db *pgxpool.Pool, login string, time time.Time) error {
	_, err := db.Exec(context.Background(), `UPDATE user_info SET last_login_time=$1 WHERE login=$2`, time, login)
	if err != nil {
		return err
	}
	return nil
}

func GetPersonData(db *pgxpool.Pool, login string) (models.PersonData, error) {
	var profile models.PersonData
	err := db.QueryRow(context.Background(),
		"SELECT name, surname, birth_date, email, phone_number FROM user_info AS t1 JOIN user_additional as t2 ON t1.user_id = t2.user_id WHERE t1.login=$1", login).
		Scan(&profile.Name, &profile.Surname, &profile.BirthDate, &profile.Email, &profile.PhoneNumber)

	if err != nil {
		return models.PersonData{}, err
	}
	return profile, nil
}

func UpdatePersonData(db *pgxpool.Pool, req models.UpdateProfileData) error {
	if len(req.Name) == 0 && len(req.Surname) == 0 && len(req.Email) == 0 && len(req.PhoneNumber) == 0 &&
		len(req.BirthDate) == 0 {
		return ErrorEmptyRequest
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
		currentBirthDate   string
		currentEmail       string
		currentPhoneNumber string
	)
	var userID string

	err := db.QueryRow(context.Background(),
		"SELECT t1.user_id, name, surname, birth_date, email, phone_number FROM user_info AS t1 JOIN user_additional as t2 ON t1.user_id = t2.user_id WHERE t1.login=$1", req.Login).
		Scan(&userID, &currentName, &currentSurname, &currentBirthDate, &currentEmail, &currentPhoneNumber)

	if err != nil {
		return ErrorDataBase
	}

	if len(newName) == 0 {
		newName = currentName
	}
	if len(newSurname) == 0 {
		newSurname = currentSurname
	}
	if len(newBirthDate) == 0 {
		newBirthDate = currentBirthDate
	}
	if len(newEmail) == 0 {
		newEmail = currentEmail
	}
	if len(newPhoneNumber) == 0 {
		newPhoneNumber = currentPhoneNumber
	}

	_, err = db.Exec(context.Background(),
		"UPDATE user_additional SET name=$1, surname=$2, birth_date=$3, email=$4, phone_number=$5 WHERE user_id=$6",
		newName, newSurname, newBirthDate, newEmail, newPhoneNumber, userID)

	if err != nil {
		return ErrorDataBase
	}

	_, err = db.Exec(context.Background(),
		"UPDATE user_info SET changed_at=$1 WHERE user_id=$2",
		time.Now(), userID)

	if err != nil {
		return ErrorDataBase
	}

	return nil
}
