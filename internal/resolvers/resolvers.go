package resolvers

import (
	"context"
	"reddit/internal/entities"
	"reddit/internal/usecases"
	"strconv"
)

// Resolver - основной резолвер
type Resolver struct {
	PostUseCase    *usecases.PostUseCase
	CommentUseCase *usecases.CommentUseCase
}

func (r *Resolver) Comment() CommentResolver {
	return &commentResolver{r}
}

type commentResolver struct{ *Resolver }

// Реализация метода `ID`
func (r *commentResolver) ID(ctx context.Context, obj *entities.Comment) (string, error) {
	return strconv.Itoa(obj.ID), nil // Преобразуем ID в строку
}

// Реализация метода `PostID`
func (r *commentResolver) PostID(ctx context.Context, obj *entities.Comment) (string, error) {
	return strconv.Itoa(obj.PostID), nil // Преобразуем PostID в строку
}

// Реализация метода `ParentID`
func (r *commentResolver) ParentID(ctx context.Context, obj *entities.Comment) (*string, error) {
	if obj.ParentID != nil {
		parentIDStr := strconv.Itoa(*obj.ParentID)
		return &parentIDStr, nil
	}
	return nil, nil // Если нет родительского комментария, возвращаем nil
}
func (r *Resolver) Post() PostResolver {
	return &postResolver{r} // Возвращаем `postResolver`
}

type postResolver struct{ *Resolver }

// Реализация метода `ID`
func (r *postResolver) ID(ctx context.Context, obj *entities.Post) (string, error) {
	return strconv.Itoa(obj.ID), nil // Преобразуем ID в строку
}

// MutationResolver - реализация резолвера для мутаций
type mutationResolver struct{ *Resolver }

// QueryResolver - реализация резолвера для запросов
type queryResolver struct{ *Resolver }

// Реализация метода для создания нового поста
func (r *mutationResolver) CreatePost(ctx context.Context, title string, content string, allowComments bool) (*entities.Post, error) {
	return r.PostUseCase.CreateNewPost(title, content, allowComments)
}

// Реализация метода для создания нового комментария
func (r *mutationResolver) CreateComment(ctx context.Context, postID string, content string, parentID *string) (*entities.Comment, error) {
	// Преобразуем строковый идентификатор в целое число
	id, err := strconv.Atoi(postID)
	if err != nil {
		return nil, err
	}

	var pID *int
	if parentID != nil {
		parentInt, err := strconv.Atoi(*parentID)
		if err != nil {
			return nil, err
		}
		pID = &parentInt
	}

	return r.CommentUseCase.CreateComment(id, content, pID)
}

// Реализация метода для получения всех постов
func (r *queryResolver) Posts(ctx context.Context) ([]*entities.Post, error) {
	return r.PostUseCase.GetAllPosts()
}

// Реализация метода для получения поста по идентификатору
func (r *queryResolver) Post(ctx context.Context, id string) (*entities.Post, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	return r.PostUseCase.GetPost(idInt)
}

// Реализация метода для получения всех комментариев к посту
func (r *Resolver) Comments(ctx context.Context, PostID string, ParentId *string, Offset *int, Limit *int) ([]*entities.Comment, error) {
	postIDInt, err := strconv.Atoi(PostID)
	if err != nil {
		return nil, err
	}
	parentId, err := strconv.Atoi(*ParentId)
	return r.CommentUseCase.GetComments(postIDInt, &parentId, *Offset, *Limit)
}

// Mutation возвращает MutationResolver
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

// Query возвращает QueryResolver
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
