package models

import "time"

type RegisterRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type GetPersonDataRequest struct {
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	BirthDate   time.Time `json:"birth_date,omitempty"`
	Email       string    `json:"email,omitempty"`
	PhoneNumber string    `json:"phone_number,omitempty"`
}

type UpdatePersonDataRequest struct {
	Name        string    `json:"name,omitempty"`
	Surname     string    `json:"surname,omitempty"`
	BirthDate   time.Time `json:"birth_date,omitempty"`
	Email       string    `json:"email,omitempty"`
	PhoneNumber string    `json:"phone_number,omitempty"`
}
