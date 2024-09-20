package services

import (
	"context"
	"fmt"
	"golang-backend-test/app/commons/response"
	"golang-backend-test/app/models"
	"golang-backend-test/app/params"
	"golang-backend-test/app/repositories"
	"time"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type AuthorService interface {
	FindDetailAuthor(ctx context.Context, id int) (*params.AuthorResponse, *response.CustomError)
	FindAllAuthors(ctx context.Context) ([]*params.AuthorResponse, *response.CustomError)
	CrateAuthor(ctx context.Context, req *params.AuthorRequest) *response.CustomError
	UpdateAuthor(ctx context.Context, id int, req *params.AuthorRequest) (*params.AuthorResponse, *response.CustomError)
	DeleteAuthor(ctx context.Context, id int) *response.CustomError
}

type AuthorServiceImpl struct {
	AuthorRepository repositories.AuthorRepository
	DB               *gorm.DB
}

func NewAuthorService(authorRepository repositories.AuthorRepository, db *gorm.DB) AuthorService {
	return &AuthorServiceImpl{
		AuthorRepository: authorRepository,
		DB:               db,
	}
}

func (service *AuthorServiceImpl) FindDetailAuthor(ctx context.Context, id int) (*params.AuthorResponse, *response.CustomError) {
	author, err := service.AuthorRepository.FindAuthorById(ctx, service.DB, id)
	if err != nil {
		return nil, response.NotFoundError()
	}

	return &params.AuthorResponse{
		ID:        author.ID,
		Name:      author.Name,
		Birthdate: author.Birthdate.Format("2006-01-02"),
	}, nil

}

func (service *AuthorServiceImpl) FindAllAuthors(ctx context.Context) ([]*params.AuthorResponse, *response.CustomError) {
	authors, err := service.AuthorRepository.GetListAuthors(ctx, service.DB)
	if err != nil {
		return nil, response.BadRequestError()
	}
	var AuthorResponses []*params.AuthorResponse
	for _, author := range authors {
		AuthorResponses = append(AuthorResponses, &params.AuthorResponse{
			ID:        author.ID,
			Name:      author.Name,
			Birthdate: author.Birthdate.Format("2006-01-02"),
		})
	}
	return AuthorResponses, nil
}

func (service *AuthorServiceImpl) CrateAuthor(ctx context.Context, req *params.AuthorRequest) *response.CustomError {
	val := validator.New()
	err := val.Struct(req)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errors []interface{}
		for _, fieldError := range validationErrors {
			error := "error " + fieldError.Field() + " on tag " + fieldError.Tag()
			errors = append(errors, error)
		}
		return response.BadRequestErrorWithAdditionalInfo(errors)
	}

	var author = new(models.Author)
	author.Name = req.Name
	birthdate, err := time.Parse("2006-01-02", req.Birthdate)
	if err != nil {
		return response.BadRequestErrorWithAdditionalInfo(fmt.Sprintf("Invalid date format: %s", err.Error()))
	}
	author.Birthdate = birthdate
	if err := service.AuthorRepository.CreateAuthor(ctx, service.DB, author); err != nil {
		return response.BadRequestError()
	}

	return nil
}

func (service *AuthorServiceImpl) UpdateAuthor(ctx context.Context, id int, req *params.AuthorRequest) (*params.AuthorResponse, *response.CustomError) {
	val := validator.New()
	err := val.Struct(req)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errors []interface{}
		for _, fieldError := range validationErrors {
			error := "error " + fieldError.Field() + " on tag " + fieldError.Tag()
			errors = append(errors, error)
		}
		return nil, response.BadRequestErrorWithAdditionalInfo(errors)
	}

	var author = new(models.Author)
	author.ID = uint(id)
	author.Name = req.Name
	birthdate, _ := time.Parse("2006-01-02", req.Birthdate)
	// if err != nil {
	// 	return nil, response.BadRequestErrorWithAdditionalInfo(fmt.Sprintf("Invalid date format: %s", err.Error()))
	// }
	author.Birthdate = birthdate
	if err := service.AuthorRepository.UpdateAuthor(ctx, service.DB, author); err != nil {
		return nil, response.BadRequestError()
	}

	return &params.AuthorResponse{
		ID:        author.ID,
		Name:      author.Name,
		Birthdate: author.Birthdate.Format("2006-01-02"),
	}, nil
}

func (service *AuthorServiceImpl) DeleteAuthor(ctx context.Context, id int) *response.CustomError {
	err := service.AuthorRepository.DeleteAuthor(ctx, service.DB, id)
	if err != nil {
		return response.NotFoundError()
	}

	return nil
}
