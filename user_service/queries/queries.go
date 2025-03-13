package queries

import (
	"context"
	"errors"
	"log"
	"time"
	"user_service/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterUser(db *pgxpool.Pool, req *models.RegistrationData, hashedPassword string) error {
	queryUserInfo := `
		INSERT INTO user_info (user_id, login, password_hash, registered_at, changed_at, last_login_time) 
		VALUES ($1, $2, $3, $4, $5, $5, $5)`
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
