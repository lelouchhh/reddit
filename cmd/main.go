package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"reddit/internal/infrastructure/database"
	"reddit/internal/infrastructure/server"
	"reddit/internal/repositories/memory"
	"reddit/pkg/logger"

	"reddit/internal/repositories/postgres"
	"reddit/internal/resolvers"
	"reddit/internal/usecases"
)

func main() {
	// Загрузка переменных окружения из .env файла
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("No .env file found: %v", err) // Сообщение, если .env файл не найден
	}
	logger.GlobalLogger.Info("Application is starting")
	// Получение типа хранилища из переменной окружения
	storageType := os.Getenv("STORAGE")
	fmt.Println(storageType)
	var postRepo postgres.PostRepository
	var commentRepo postgres.CommentRepository
	// Инициализация репозиториев на основе типа хранилища
	if storageType == "postgres" {
		db, err := database.NewPostgres()
		if err != nil {
			logger.GlobalLogger.Error(err.Error())
			log.Fatalf("Error connecting to PostgreSQL: %v", err)
		}
		postRepo = postgres.NewPostgresPostRepository(db)
		commentRepo = postgres.NewPostgresCommentRepository(db)
	} else if storageType == "memory" {
		postRepo = memory.NewInMemoryPostRepository()
		commentRepo = memory.NewInMemoryCommentRepository()
	} else {
		logger.GlobalLogger.Error("can't choose storage type")

		log.Fatalf("Unknown storage type: %s", storageType)
	}

	// Инициализация сценариев использования
	postUseCase := usecases.NewPostUseCase(postRepo)
	commentUseCase := usecases.NewCommentUseCase(commentRepo)

	// Инициализация резолверов
	schema := resolvers.NewExecutableSchema(resolvers.Config{
		Resolvers: &resolvers.Resolver{
			PostUseCase:    postUseCase,
			CommentUseCase: commentUseCase,
		},
	})

	// Запуск HTTP-сервера для GraphQL
	httpServer := server.NewHTTPServer(schema)

	// Запуск HTTP-сервера
	logger.GlobalLogger.Info("Starting server on port 8080")
	if err := httpServer.ListenAndServe(); err != nil {
		logger.GlobalLogger.Error(err.Error())

		log.Fatalf("Failed to start server: %v", err)
	}
}
