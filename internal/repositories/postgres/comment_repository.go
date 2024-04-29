package postgres

import (
	"database/sql"
	"reddit/internal/entities"
)

// CommentRepository - интерфейс для репозитория комментариев
type CommentRepository interface {
	CreateComment(postID int, content string, parentID *int) (*entities.Comment, error)    // Создает новый комментарий
	GetComments(postID int, parentId *int, offset, limit int) ([]*entities.Comment, error) // Возвращает список комментариев для поста
}

// PostgresCommentRepository - реализация репозитория комментариев для PostgreSQL
type PostgresCommentRepository struct {
	db *sql.DB // Объект базы данных
}

// NewPostgresCommentRepository - конструктор для создания нового репозитория
func NewPostgresCommentRepository(db *sql.DB) *PostgresCommentRepository {
	return &PostgresCommentRepository{db: db}
}

func (repo *PostgresCommentRepository) CreateComment(postID int, content string, parentID *int) (*entities.Comment, error) {
	// Добавление нового комментария
	var commentID int
	err := repo.db.QueryRow(
		"INSERT INTO comments (post_id, content, parent_id) VALUES ($1, $2, $3) RETURNING id",
		postID, content, parentID,
	).Scan(&commentID)
	if err != nil {
		return nil, err
	}

	return &entities.Comment{
		ID:       commentID,
		PostID:   postID,
		Content:  content,
		ParentID: parentID,
	}, nil
}

func (repo *PostgresCommentRepository) GetComments(postID int, parentId *int, offset int, limit int) ([]*entities.Comment, error) {
	query := "SELECT id, post_id, content, parent_id FROM comments WHERE post_id = $1 ORDER BY id LIMIT $2 OFFSET $3"
	rows, err := repo.db.Query(query, postID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*entities.Comment
	for rows.Next() {
		var comment entities.Comment
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.Content, &comment.ParentID); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}

	return comments, nil
}
