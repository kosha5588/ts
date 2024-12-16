package server

import (
	"net/http"
	"ts/backend/internal/handlers"
)

type Server struct {
	port string
}

// NewServer создает новый сервер с указанным портом
func NewServer(port string) *Server {
	return &Server{port: port}
}

// ListenAndServe запускает сервер
func (s *Server) ListenAndServe() error {
	// Указываем путь к директории с публичными файлами
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("../../Public/assets/"))))

	// Обработчик для шаблонов
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/reg", handlers.RegHandler)
	http.HandleFunc("/signIn", handlers.SignInHandler)

	// Запускаем HTTP сервер
	return http.ListenAndServe(s.port, nil)
}
