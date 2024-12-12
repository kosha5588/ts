package server

import (
	"net/http"
	"path/filepath"
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
	publicDir := filepath.Join("..", "..", "Public")

	// Обработчик для статических файлов
	fs := http.FileServer(http.Dir(publicDir))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Обработчик для шаблонов
	http.HandleFunc("/", handlers.IndexHandler)

	// Запускаем HTTP сервер
	return http.ListenAndServe(s.port, nil)
}
