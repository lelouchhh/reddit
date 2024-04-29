package memory

import (
	"reddit/internal/entities"
)

// InMemoryPostRepository - репозиторий постов, хранящий данные в памяти
type InMemoryPostRepository struct {
	posts  map[int]*entities.Post // Хранилище постов в памяти
	nextID int                    // Следующий доступный идентификатор
}

func (repo *InMemoryPostRepository) GetPostByID(id int) (*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}

// Создание нового репозитория постов в памяти
func NewInMemoryPostRepository() *InMemoryPostRepository {
	return &InMemoryPostRepository{
		posts:  make(map[int]*entities.Post),
		nextID: 1,
	}
}

// Создание нового поста
func (repo *InMemoryPostRepository) CreatePost(title, content string, allowComments bool) (*entities.Post, error) {
	id := repo.nextID
	repo.nextID++
	post := &entities.Post{
		ID:            id,
		Title:         title,
		Content:       content,
		AllowComments: allowComments,
	}
	repo.posts[id] = post
	return post, nil
}

// Получение всех постов
func (repo *InMemoryPostRepository) GetPosts() ([]*entities.Post, error) {
	var posts []*entities.Post
	for _, post := range repo.posts {
		posts = append(posts, post)
	}
	return posts, nil
}
