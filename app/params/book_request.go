package params

type BookRequest struct {
	Title    string `json:"title" validate:"required"`
	ISBN     string `json:"isbn"`
	AuthorID uint   `json:"author_id" validate:"required"`
}
