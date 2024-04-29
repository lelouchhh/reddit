package usecases

import (
	"database/sql"
	"fmt"
	"reddit/internal/entities"
	"reddit/internal/repositories/postgres"
)

// PostUseCase - сценарий использования для работы с постами
type PostUseCase struct {
	repo postgres.PostRepository
}

// NewPostUseCase - создает новый экземпляр `PostUseCase`
func NewPostUseCase(repo postgres.PostRepository) *PostUseCase {
	return &PostUseCase{repo: repo}
}

// GetAllPosts возвращает список всех постов
func (uc *PostUseCase) GetAllPosts() ([]*entities.Post, error) {
	return uc.repo.GetPosts() // Используем метод репозитория для получения всех постов
}

// GetPost возвращает пост по его идентификатору
func (uc *PostUseCase) GetPost(id int) (*entities.Post, error) {
	post, err := uc.repo.GetPostByID(id) // Используем репозиторий для получения поста по идентификатору
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("post with ID %d not found", id) // Если пост не найден, возвращаем ошибку
		}
		return nil, err // Возвращаем любую другую ошибку
	}
	return post, nil
}

// CreateNewPost создает новый пост
func (uc *PostUseCase) CreateNewPost(title, content string, allowComments bool) (*entities.Post, error) {
	if title == "" {
		return nil, fmt.Errorf("title cannot be empty") // Проверка, что заголовок не пустой
	}

	return uc.repo.CreatePost(title, content, allowComments) // Используем репозиторий для создания поста
}
