package entities

type Comment struct {
	ID       int    // Идентификатор комментария
	PostID   int    // Идентификатор связанного поста
	Content  string // Содержание комментария
	ParentID *int   // Идентификатор родительского комментария (если есть)
	Children []*Comment
}
