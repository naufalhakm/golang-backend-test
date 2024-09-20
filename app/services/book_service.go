package services

import (
	"context"
	"golang-backend-test/app/commons/response"
	"golang-backend-test/app/models"
	"golang-backend-test/app/params"
	"golang-backend-test/app/repositories"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type BookService interface {
	FindDetailBook(ctx context.Context, id int) (*params.BookResponse, *response.CustomError)
	FindAllBooks(ctx context.Context) ([]*params.BookResponse, *response.CustomError)
	CrateBook(ctx context.Context, req *params.BookRequest) *response.CustomError
	UpdateBook(ctx context.Context, id int, req *params.BookRequest) (*params.BookResponse, *response.CustomError)
	DeleteBook(ctx context.Context, id int) *response.CustomError
}

type BookServiceImpl struct {
	BookRepository   repositories.BookRepository
	AuthorRepository repositories.AuthorRepository
	DB               *gorm.DB
}

func NewBookService(bookRepository repositories.BookRepository, authorRepository repositories.AuthorRepository, db *gorm.DB) BookService {
	return &BookServiceImpl{
		BookRepository:   bookRepository,
		AuthorRepository: authorRepository,
		DB:               db,
	}
}

func (service *BookServiceImpl) FindDetailBook(ctx context.Context, id int) (*params.BookResponse, *response.CustomError) {
	book, err := service.BookRepository.FindBookById(ctx, service.DB, id)
	if err != nil {
		return nil, response.NotFoundError()
	}

	return &params.BookResponse{
		ID:    book.ID,
		Title: book.Title,
		ISBN:  book.ISBN,
		AuthorResponse: &params.AuthorResponse{
			ID:        book.AuthorID,
			Name:      book.Author.Name,
			Birthdate: book.Author.Birthdate.Format("2006-01-02"),
		},
	}, nil

}

func (service *BookServiceImpl) FindAllBooks(ctx context.Context) ([]*params.BookResponse, *response.CustomError) {
	books, err := service.BookRepository.GetListBooks(ctx, service.DB)
	if err != nil {
		return nil, response.BadRequestError()
	}
	var bookResponses []*params.BookResponse
	for _, book := range books {
		bookResponses = append(bookResponses, &params.BookResponse{
			ID:    book.ID,
			Title: book.Title,
			ISBN:  book.ISBN,
			AuthorResponse: &params.AuthorResponse{
				ID:        book.AuthorID,
				Name:      book.Author.Name,
				Birthdate: book.Author.Birthdate.Format("2006-01-02"),
			},
		})
	}
	return bookResponses, nil
}

func (service *BookServiceImpl) CrateBook(ctx context.Context, req *params.BookRequest) *response.CustomError {
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

	var book = new(models.Book)
	book.Title = req.Title
	book.ISBN = req.ISBN
	book.AuthorID = req.AuthorID
	if err := service.BookRepository.CreateBook(ctx, service.DB, book); err != nil {
		return response.BadRequestError()
	}

	return nil
}

func (service *BookServiceImpl) UpdateBook(ctx context.Context, id int, req *params.BookRequest) (*params.BookResponse, *response.CustomError) {
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

	newAuthor, err := service.AuthorRepository.FindAuthorById(ctx, service.DB, int(req.AuthorID))
	if err != nil {
		return nil, response.BadRequestError()
	}

	var book = new(models.Book)
	book.ID = uint(id)
	book.Title = req.Title
	book.ISBN = req.ISBN
	book.AuthorID = newAuthor.ID

	if err := service.BookRepository.UpdateBook(ctx, service.DB, book); err != nil {
		return nil, response.BadRequestError()
	}

	return &params.BookResponse{
		ID:    book.ID,
		Title: book.Title,
		ISBN:  book.ISBN,
		AuthorResponse: &params.AuthorResponse{
			ID:        newAuthor.ID,
			Name:      newAuthor.Name,
			Birthdate: newAuthor.Birthdate.Format("2006-01-02"),
		},
	}, nil
}

func (service *BookServiceImpl) DeleteBook(ctx context.Context, id int) *response.CustomError {
	err := service.BookRepository.DeleteBook(ctx, service.DB, id)
	if err != nil {
		return response.NotFoundError()
	}

	return nil
}
