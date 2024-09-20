package params

type AuthorRequest struct {
	Name      string `json:"name" validate:"required"`
	Birthdate string `json:"birthdate" validate:"required"`
}
