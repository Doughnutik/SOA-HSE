package models

type RegistrationData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type AuthData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type PersonData struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	BirthDate   string `json:"birth_date"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type UpdateProfileData struct {
	Login       string `json:"login"`
	Password    string `json:"password"`
	Name        string `json:"name,omitempty"`
	Surname     string `json:"surname,omitempty"`
	BirthDate   string `json:"birth_date,omitempty"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}
