package server

import (
	"api_gateway/config"
	"api_gateway/post_service_api"
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
	s.Router.Post("/login", user_service_api.LoginHandler)
	s.Router.Get("/profile", user_service_api.GetProfile)
	s.Router.Put("/profile", user_service_api.UpdateProfile)
	s.Router.Post("/posts", post_service_api.CreatePostHandler) // Проксирование запроса на post_service
	s.Router.Get("/posts/{post_id}", post_service_api.GetPostHandler)
	s.Router.Put("/posts/{post_id}", post_service_api.UpdatePostHandler)
	s.Router.Delete("/posts/{post_id}", post_service_api.DeletePostHandler)
	s.Router.Get("/posts/list", post_service_api.ListPostsHandler)
}
