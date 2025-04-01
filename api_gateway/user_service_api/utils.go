package user_service_api

import (
	"api_gateway/config"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

var (
	ErrorIncorrectRequest = errors.New("некорректные параметры запроса")
	ErrorInternal         = errors.New("внутренняя ошибка сервиса")
	Cfg                   config.Config
)

func validateRequestBody[T any](r *http.Request, v *T) error {
	if r.Header.Get("Content-Type") != "application/json" {
		return ErrorIncorrectRequest
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("validateRequestBody\t Ошибка чтения запроса: %v", err)
		return ErrorInternal
	}
	r.Body = io.NopCloser(bytes.NewBuffer(body))
	if err := json.Unmarshal(body, v); err != nil {
		return ErrorIncorrectRequest
	}
	return nil
}

func proxyRequest(w http.ResponseWriter, r *http.Request, endpoint string) {
	url := Cfg.UserServiceUrl + endpoint

	// Создаем новый HTTP-запрос с оригинальным методом
	req, err := http.NewRequest(r.Method, url, r.Body)
	if err != nil {
		http.Error(w, "ошибка создания запроса к user_service", http.StatusInternalServerError)
		log.Printf("proxyRequest\t ошибка создания запроса: %v", err)
		return
	}

	// Копируем заголовки
	for key, values := range r.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	// Используем http.Client для отправки запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "ошибка запроса к user_service", http.StatusInternalServerError)
		log.Printf("proxyRequest\t ошибка запроса к user_service: %v", err)
		return
	}
	defer resp.Body.Close()

	// Копируем ответ user_service в клиентский ответ
	copyResponse(w, resp)
}

func copyResponse(w http.ResponseWriter, resp *http.Response) {
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
