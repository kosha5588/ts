package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func DatabaseTs() {
	connStr := "host=db user=user password=password dbname=db sslmode=disable" // Используйте имя сервиса
	var db *sql.DB
	var err error

	// Попробуйте подключиться несколько раз
	for i := 0; i < 10; i++ { // Увеличьте количество попыток
		db, err = sql.Open("postgres", connStr)
		if err == nil {
			break
		}
		log.Printf("Ошибка подключения: %v. Повторная попытка через 5 секунд...", err) // Увеличьте время ожидания
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}
	defer db.Close()

	// Проверка подключения
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Успешно подключено к базе данных!")

	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name TEXT,
        age INT
    );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Таблица успешно создана!")
}
