package services

import (
	"context"
	"errors"
	"golang-backend-test/app/models"
	"golang-backend-test/app/params"
	"golang-backend-test/app/repositories"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestFindDetailBook_Success(t *testing.T) {
	bookRepo := new(repositories.MockBookRepository)
	authorRepo := new(repositories.MockAuthorRepository)
	db := new(gorm.DB)
	bookID := uint(1)
	authorID := uint(1)

	book := &models.Book{
		ID:       bookID,
		Title:    "Test Book",
		ISBN:     "123456789",
		AuthorID: authorID,
		Author: models.Author{
			ID:        authorID,
			Name:      "Test Author",
			Birthdate: time.Date(1985, time.April, 5, 0, 0, 0, 0, time.UTC),
		},
	}

	bookRepo.On("FindBookById", mock.Anything, db, int(bookID)).Return(book, nil)
	service := NewBookService(bookRepo, authorRepo, db)

	result, err := service.FindDetailBook(context.Background(), int(bookID))

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, bookID, result.ID)
	assert.Equal(t, "Test Book", result.Title)
	assert.Equal(t, "123456789", result.ISBN)
	assert.Equal(t, "Test Author", result.AuthorResponse.Name)
	assert.Equal(t, "1985-04-05", result.AuthorResponse.Birthdate)

	bookRepo.AssertExpectations(t)
}

func TestFindDetailBook_BookNotFound(t *testing.T) {
	bookRepo := new(repositories.MockBookRepository)
	authorRepo := new(repositories.MockAuthorRepository)
	db := new(gorm.DB)
	bookID := uint(1)

	bookRepo.On("FindBookById", mock.Anything, db, int(bookID)).Return(nil, errors.New("book not found"))
	service := NewBookService(bookRepo, authorRepo, db)

	result, err := service.FindDetailBook(context.Background(), int(bookID))

	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "NOT FOUND ERROR", err.Message)
	assert.Equal(t, 400, err.StatusCode)

	bookRepo.AssertExpectations(t)
}

func TestFindAllBook_Success(t *testing.T) {
	bookRepo := new(repositories.MockBookRepository)
	authorRepo := new(repositories.MockAuthorRepository)
	db := new(gorm.DB)

	books := []*models.Book{
		{
			ID:       1,
			Title:    "Test Book",
			ISBN:     "123456789",
			AuthorID: 1,
			Author: models.Author{
				ID:        1,
				Name:      "Test Author",
				Birthdate: time.Date(1985, time.April, 5, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	bookRepo.On("GetListBooks", mock.Anything, db).Return(books, nil)
	service := NewBookService(bookRepo, authorRepo, db)

	result, err := service.FindAllBooks(context.Background())

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, "Test Book", result[0].Title)
	assert.Equal(t, "123456789", result[0].ISBN)
	assert.Equal(t, "Test Author", result[0].AuthorResponse.Name)
	assert.Equal(t, "1985-04-05", result[0].AuthorResponse.Birthdate)

	bookRepo.AssertExpectations(t)
}

func TestFindAllBooks_RepositoryError(t *testing.T) {
	bookRepo := new(repositories.MockBookRepository)
	authorRepo := new(repositories.MockAuthorRepository)
	db := new(gorm.DB)
	service := NewBookService(bookRepo, authorRepo, db)

	bookRepo.On("GetListBooks", mock.Anything, db).Return(nil, errors.New("db error"))

	result, err := service.FindAllBooks(context.Background())

	assert.NotNil(t, err)
	assert.Nil(t, result)
	bookRepo.AssertExpectations(t)
}

func TestCreateBook_Success(t *testing.T) {
	bookRepo := new(repositories.MockBookRepository)
	authorRepo := new(repositories.MockAuthorRepository)
	db := new(gorm.DB)
	service := NewBookService(bookRepo, authorRepo, db)

	validRequest := &params.BookRequest{
		Title:    "Test Book",
		ISBN:     "123456789",
		AuthorID: 1,
	}

	bookRepo.On("CreateBook", mock.Anything, db, mock.AnythingOfType("*models.Book")).Return(nil)

	err := service.CrateBook(context.Background(), validRequest)

	assert.Nil(t, err)
	bookRepo.AssertExpectations(t)
}

func TestCreateBook_ValidationError(t *testing.T) {
	bookRepo := new(repositories.MockBookRepository)
	authorRepo := new(repositories.MockAuthorRepository)
	db := new(gorm.DB)
	service := NewBookService(bookRepo, authorRepo, db)

	invalidRequest := &params.BookRequest{
		Title: "",
		ISBN:  "123456789",
	}

	err := service.CrateBook(context.Background(), invalidRequest)

	assert.NotNil(t, err)
	assert.Equal(t, "BAD REQUEST ERROR", err.Message)
}

func TestCreateBook_RepositoryError(t *testing.T) {
	bookRepo := new(repositories.MockBookRepository)
	authorRepo := new(repositories.MockAuthorRepository)
	db := new(gorm.DB)
	service := NewBookService(bookRepo, authorRepo, db)

	validRequest := &params.BookRequest{
		Title:    "Test Book",
		ISBN:     "123456789",
		AuthorID: 1,
	}

	bookRepo.On("CreateBook", mock.Anything, db, mock.AnythingOfType("*models.Book")).Return(errors.New("db error"))

	err := service.CrateBook(context.Background(), validRequest)

	assert.NotNil(t, err)
	assert.Equal(t, "BAD REQUEST ERROR", err.Message)
	bookRepo.AssertExpectations(t)
}

func TestUpdateBook_Success(t *testing.T) {
	bookRepo := new(repositories.MockBookRepository)
	authorRepo := new(repositories.MockAuthorRepository)
	db := new(gorm.DB)
	service := NewBookService(bookRepo, authorRepo, db)

	validRequest := &params.BookRequest{
		Title:    "Updated Book",
		ISBN:     "123456789",
		AuthorID: 1,
	}

	bookRepo.On("UpdateBook", mock.Anything, db, mock.AnythingOfType("*models.Book")).Return(nil)
	authorRepo.On("FindAuthorById", mock.Anything, db, 1).Return(&models.Author{
		ID:        1,
		Name:      "Author 1",
		Birthdate: time.Date(1985, time.April, 5, 0, 0, 0, 0, time.UTC),
	}, nil)

	result, err := service.UpdateBook(context.Background(), 1, validRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Updated Book", result.Title)
	bookRepo.AssertExpectations(t)
	authorRepo.AssertExpectations(t)
}

func TestUpdateBook_ValidationError(t *testing.T) {
	bookRepo := new(repositories.MockBookRepository)
	authorRepo := new(repositories.MockAuthorRepository)
	db := new(gorm.DB)
	service := NewBookService(bookRepo, authorRepo, db)

	invalidRequest := &params.BookRequest{
		Title: "",
		ISBN:  "123456789",
	}

	result, err := service.UpdateBook(context.Background(), 1, invalidRequest)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "BAD REQUEST ERROR", err.Message)
}

func TestUpdateBook_AuthorNotFound(t *testing.T) {
	bookRepo := new(repositories.MockBookRepository)
	authorRepo := new(repositories.MockAuthorRepository)
	db := new(gorm.DB)
	service := NewBookService(bookRepo, authorRepo, db)

	invalidRequest := &params.BookRequest{
		Title:    "Update Book",
		ISBN:     "123456789",
		AuthorID: 3,
	}

	authorRepo.On("FindAuthorById", mock.Anything, db, 3).Return(nil, errors.New("author not found"))

	result, err := service.UpdateBook(context.Background(), 1, invalidRequest)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "BAD REQUEST ERROR", err.Message)
	authorRepo.AssertExpectations(t)
}

func TestUpdateBook_RepositoryError(t *testing.T) {
	bookRepo := new(repositories.MockBookRepository)
	authorRepo := new(repositories.MockAuthorRepository)
	db := new(gorm.DB)
	service := NewBookService(bookRepo, authorRepo, db)

	validRequest := &params.BookRequest{
		Title:    "Updated Book",
		ISBN:     "123456789",
		AuthorID: 1,
	}

	authorRepo.On("FindAuthorById", mock.Anything, db, 1).Return(&models.Author{
		ID:        1,
		Name:      "Test Author",
		Birthdate: time.Date(1985, time.April, 5, 0, 0, 0, 0, time.UTC),
	}, nil)

	bookRepo.On("UpdateBook", mock.Anything, db, mock.AnythingOfType("*models.Book")).Return(errors.New("db error"))

	result, err := service.UpdateBook(context.Background(), 1, validRequest)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "BAD REQUEST ERROR", err.Message)
	bookRepo.AssertExpectations(t)
}

func TestDeleteBook_Success(t *testing.T) {
	bookRepo := new(repositories.MockBookRepository)
	authorRepo := new(repositories.MockAuthorRepository)
	db := new(gorm.DB)
	service := NewBookService(bookRepo, authorRepo, db)

	bookRepo.On("DeleteBook", mock.Anything, db, 1).Return(nil)

	err := service.DeleteBook(context.Background(), 1)

	assert.Nil(t, err)
	bookRepo.AssertExpectations(t)
}

func TestDeleteBook_RepositoryError(t *testing.T) {
	bookRepo := new(repositories.MockBookRepository)
	authorRepo := new(repositories.MockAuthorRepository)
	db := new(gorm.DB)
	service := NewBookService(bookRepo, authorRepo, db)

	bookRepo.On("DeleteBook", mock.Anything, db, 1).Return(errors.New("book not found"))

	err := service.DeleteBook(context.Background(), 1)

	assert.NotNil(t, err)
	assert.Equal(t, "NOT FOUND ERROR", err.Message)
	bookRepo.AssertExpectations(t)
}
