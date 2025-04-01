package main

import (
	"api_gateway/config"
	"api_gateway/post_service_api"
	"api_gateway/server"
	"api_gateway/user_service_api"
	"log"
	"net/http"
	"os"
)

func main() {
	logFile, err := os.OpenFile("api_gateway.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Ошибка открытия log файла: %v", err)
	}
	log.SetOutput(logFile)
	log.Print("\n\n")

	// Загружаем конфигурацию
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Ошибка загрузки конфигурации")
	}
	user_service_api.Cfg = *cfg
	post_service_api.Cfg = *cfg

	srv := server.NewServer(cfg)
	log.Printf("Сервер запущен на порту %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, srv.Router))
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/go-chi/chi"
// )

// func GetPostHandler(w http.ResponseWriter, r *http.Request) {
// 	// Извлекаем post_id из URL
// 	postID := chi.URLParam(r, "post_id")
// 	log.Printf("Получен post_id: %s", postID)

// 	// Возвращаем post_id в ответе
// 	fmt.Fprintf(w, "Post ID: %s", postID)
// }

// func main() {
// 	r := chi.NewRouter()

// 	// Определение маршрута с параметром post_id
// 	r.Get("/posts/{post_id}", GetPostHandler)

// 	log.Println("Сервер слушает на порту 8081")
// 	log.Fatal(http.ListenAndServe(":8081", r))
// }
