package postgres

import (
	"database/sql"
	"fmt"
	"reddit/internal/entities"
)

type PostRepository interface {
	GetPosts() ([]*entities.Post, error)                                          // Возвращает список всех постов
	CreatePost(title, content string, allowComments bool) (*entities.Post, error) // Создает новый пост
	GetPostByID(id int) (*entities.Post, error)                                   // Возвращает пост по его идентификатору
}
type PostgresPostRepository struct {
	db *sql.DB
}

func NewPostgresPostRepository(db *sql.DB) *PostgresPostRepository {
	return &PostgresPostRepository{db: db}
}

// GetPosts возвращает список всех постов
func (repo *PostgresPostRepository) GetPosts() ([]*entities.Post, error) {
	rows, err := repo.db.Query("SELECT id, title, content, allow_comments FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*entities.Post
	for rows.Next() {
		var post entities.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AllowComments); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return posts, nil
}

// GetPostByID возвращает пост по его идентификатору
func (repo *PostgresPostRepository) GetPostByID(id int) (*entities.Post, error) {
	var post entities.Post
	query := "SELECT id, title, content, allow_comments FROM posts WHERE id = $1"
	err := repo.db.QueryRow(query, id).Scan(&post.ID, &post.Title, &post.Content, &post.AllowComments)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("post with ID %d not found", id)
		}
		return nil, err
	}
	return &post, nil
}

// CreatePost создает новый пост
func (repo *PostgresPostRepository) CreatePost(title, content string, allowComments bool) (*entities.Post, error) {
	query := "INSERT INTO posts (title, content, allow_comments) VALUES ($1, $2, $3) RETURNING id"
	var id int
	err := repo.db.QueryRow(query, title, content, allowComments).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &entities.Post{ID: id, Title: title, Content: content, AllowComments: allowComments}, nil
}
