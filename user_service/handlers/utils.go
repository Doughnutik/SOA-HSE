package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

// readRequestBody читает тело запроса и возвращает его в виде массива байтов.
func readRequestBody(r *http.Request) ([]byte, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return body, nil
}

// writeJSONResponse отправляет JSON-ответ клиенту.
func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Ошибка сериализации ответа", http.StatusInternalServerError)
		return
	}

	w.Write(response)
}
