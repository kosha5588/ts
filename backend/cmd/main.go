package main

import (
	"log"
	"ts/backend/internal/server"
)

func main() {
	// Создаем и запускаем сервер
	srv := server.NewServer(":8088")
	log.Printf("Сервис запущен на http://localhost%s\n", ":8088")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
