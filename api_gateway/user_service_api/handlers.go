package user_service_api

import (
	"errors"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req RegistrationData
	err := validateRequestBody(r, &req)
	if errors.Is(err, ErrorIncorrectRequest) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if errors.Is(err, ErrorInternal) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	proxyRequest(w, r, "/register")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req AuthData
	err := validateRequestBody(r, &req)
	if errors.Is(err, ErrorIncorrectRequest) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if errors.Is(err, ErrorInternal) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	proxyRequest(w, r, "/login")
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var req UpdateProfileData
	err := validateRequestBody(r, &req)
	if errors.Is(err, ErrorIncorrectRequest) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if errors.Is(err, ErrorInternal) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	proxyRequest(w, r, "/profile")
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	var req AuthData
	err := validateRequestBody(r, &req)
	if errors.Is(err, ErrorIncorrectRequest) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if errors.Is(err, ErrorInternal) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	proxyRequest(w, r, "/profile")
}
