package main

import (
	"log"
	"ts/backend/db"
	"ts/backend/internal/server"
)

func main() {

	// Создаем и запускаем базу данных
	db.DatabaseTs()
	// Создаем и запускаем сервер
	srv := server.NewServer(":8090")
	log.Printf("Сервис запущен на http://localhost%s\n", ":8090")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
