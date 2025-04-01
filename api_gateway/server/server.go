package server

import (
	"api_gateway/config"
	"api_gateway/user_service_api"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Server — структура для хранения зависимостей сервера.
type Server struct {
	Router *chi.Mux
	Config *config.Config
}

// NewServer создает новый HTTP-сервер и настраивает маршруты.
func NewServer(cfg *config.Config) *Server {
	r := chi.NewRouter()

	// Middleware (логирование, восстановление после паники)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Инициализируем сервер
	srv := &Server{
		Router: r,
		Config: cfg,
	}

	// Настраиваем маршруты API Gateway
	srv.setupRoutes()

	return srv
}

// setupRoutes определяет API-эндпоинты и проксирует их в user_service.
func (s *Server) setupRoutes() {
	s.Router.Post("/register", user_service_api.RegisterHandler) // Проксирование запроса на user_service
	s.Router.Post("/login", user_service_api.LoginHandler)       // Проксирование запроса на user_service
	s.Router.Get("/profile", user_service_api.GetProfile)        // Проксирование запроса на user_service
	s.Router.Put("/profile", user_service_api.UpdateProfile)     // Проксирование запроса на user_service
}
