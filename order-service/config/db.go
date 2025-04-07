package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// ConnectToDatabase с попытками подключения
func ConnectToDatabase(host, port, user, password, dbname string) *sql.DB {
	var db *sql.DB
	var err error
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)

	// Попытки подключения к базе данных
	for i := 0; i < 5; i++ { // Попытаться подключиться 5 раз
		db, err = sql.Open("postgres", dsn)
		if err != nil {
			log.Printf("Failed to connect to database (attempt %d): %v", i+1, err)
		} else {
			// Проверка доступности базы данных
			err = db.Ping()
			if err == nil {
				log.Println("Successfully connected to the database")
				return db
			}
			log.Printf("Failed to ping database (attempt %d): %v", i+1, err)
		}
		// Если не удалось подключиться, подождать 5 секунд и попробовать снова
		time.Sleep(5 * time.Second)
	}
	log.Fatalf("Unable to connect to the database after 5 attempts: %v", err)
	return nil
}
