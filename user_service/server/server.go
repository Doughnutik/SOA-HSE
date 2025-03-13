package server

import (
	"user_service/config"
	"user_service/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Server — структура для хранения зависимостей сервера.
type Server struct {
	Router *chi.Mux
	DB     *pgxpool.Pool
	Config *config.Config
}

// NewServer создает новый HTTP-сервер и настраивает маршруты.
func NewServer(cfg *config.Config, db *pgxpool.Pool) *Server {
	r := chi.NewRouter()

	// Middleware (логирование, восстановление после паники)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Инициализируем сервер
	srv := &Server{
		Router: r,
		DB:     db,
		Config: cfg,
	}

	// Настраиваем маршруты API
	srv.setupRoutes()

	return srv
}

// setupRoutes определяет API-эндпоинты согласно OpenAPI-спецификации.
func (s *Server) setupRoutes() {
	s.Router.Post("/register", handlers.RegisterUser(s.DB)) // Регистрация
	s.Router.Post("/login", handlers.LoginUser(s.DB))       // Авторизация
	s.Router.Get("/profile", handlers.GetProfile(s.DB))     // Получение профиля
	s.Router.Put("/profile", handlers.UpdateProfile(s.DB))  // Обновление профиля
}
