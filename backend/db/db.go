package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Подключение к базе данных
	dsn := "user:password@tcp(127.0.0.1:5432)/dbname"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Проверка соединения
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Выполнение SQL запроса
	query := "SELECT id, name FROM users WHERE active = ?"
	rows, err := db.Query(query, true)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Обработка результатов
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	// Проверка на ошибки во время итерации
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
