package usecases

import (
	"fmt"
	"reddit/internal/entities"
	"reddit/internal/repositories/postgres"
)

// CommentUseCase - сценарий использования для работы с комментариями
type CommentUseCase struct {
	commentRepo postgres.CommentRepository // Используем интерфейс репозитория
}

// NewCommentUseCase - создает новый экземпляр `CommentUseCase`
func NewCommentUseCase(repo postgres.CommentRepository) *CommentUseCase {
	return &CommentUseCase{commentRepo: repo}
}

// CreateComment создает новый комментарий
func (uc *CommentUseCase) CreateComment(postID int, content string, parentID *int) (*entities.Comment, error) {
	if len(content) > 2000 {
		return nil, fmt.Errorf("comment content too long")
	}

	return uc.commentRepo.CreateComment(postID, content, parentID) // Используем репозиторий для создания комментария
}

// GetComments возвращает список комментариев для заданного поста
func (uc *CommentUseCase) GetComments(postID int, parentId *int, offset int, limit int) ([]*entities.Comment, error) {
	return uc.commentRepo.GetComments(postID, parentId, offset, limit) // Используем репозиторий для получения комментариев
}
