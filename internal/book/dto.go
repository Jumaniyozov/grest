package book

type CreateBookDTO struct {
	Name     string `json:"name"`
	AuthorID string `json:"author_id"`
}
