package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func NewPostgres() (*sql.DB, error) {
	// Загружаем переменные окружения из файла .env
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	// Получаем параметры подключения к PostgreSQL из переменных окружения
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDB := os.Getenv("POSTGRES_DB")
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresPort := os.Getenv("POSTGRES_PORT")

	// Проверка наличия необходимых переменных окружения
	if postgresUser == "" || postgresPassword == "" || postgresDB == "" || postgresHost == "" || postgresPort == "" {
		return nil, fmt.Errorf("Missing PostgreSQL connection details in environment variables\n")
	}

	// Составляем строку подключения к PostgreSQL
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		postgresUser, postgresPassword, postgresHost, postgresPort, postgresDB)
	// Подключение к базе данных PostgreSQL
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to PostgreSQ")
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Failed to ping PostgreSQL %s", err.Error())
	}
	return db, nil
}
