package params

type BookResponse struct {
	ID             uint            `json:"id"`
	Title          string          `json:"title"`
	ISBN           string          `json:"isbn"`
	AuthorResponse *AuthorResponse `json:"author,omitempty"`
}
