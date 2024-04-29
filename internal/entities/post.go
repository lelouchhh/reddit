package entities

type Post struct {
	ID            int    // Идентификатор поста
	Title         string // Заголовок поста
	Content       string // Содержание поста
	AllowComments bool   // Флаг, разрешающий или запрещающий комментарии
}
