package memory

import (
	"errors"
	"reddit/internal/entities"
)

// InMemoryCommentRepository - репозиторий комментариев, хранящий данные в памяти
type InMemoryCommentRepository struct {
	comments map[int]*entities.Comment // Хранилище комментариев в памяти
	nextID   int                       // Следующий доступный идентификатор
}

// Создание нового репозитория комментариев в памяти
func NewInMemoryCommentRepository() *InMemoryCommentRepository {
	return &InMemoryCommentRepository{
		comments: make(map[int]*entities.Comment),
		nextID:   1,
	}
}

// Создание нового комментария
func (repo *InMemoryCommentRepository) CreateComment(postID int, content string, parentID *int) (*entities.Comment, error) {
	id := repo.nextID
	repo.nextID++
	comment := &entities.Comment{
		ID:       id,
		PostID:   postID,
		Content:  content,
		ParentID: parentID,
	}
	repo.comments[id] = comment
	return comment, nil
}

// Получение комментариев по посту
// GetComments с поддержкой пагинации и рекурсии
func (repo *InMemoryCommentRepository) GetComments(postID int, parentId *int, offset int, limit int) ([]*entities.Comment, error) {
	var comments []*entities.Comment

	// Фильтруем комментарии по postID и parentId
	for _, comment := range repo.comments {
		if comment.PostID == postID && (parentId == nil || (comment.ParentID != nil && *comment.ParentID == *parentId)) {
			comments = append(comments, comment)
		}
	}

	// Применяем пагинацию
	if offset > len(comments) {
		return nil, errors.New("offset exceeds the number of available comments")
	}

	end := offset + limit
	if end > len(comments) {
		end = len(comments)
	}

	// Возвращаем ограниченное количество комментариев на основе offset и limit
	return comments[offset:end], nil
}
