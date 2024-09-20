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

func TestFindDetailAuthor_Success(t *testing.T) {
	authorRepo := new(repositories.MockAuthorRepository)
	db := new(gorm.DB)
	authorID := uint(1)

	Author := &models.Author{
		ID:        authorID,
		Name:      "Test Author",
		Birthdate: time.Date(1985, time.April, 5, 0, 0, 0, 0, time.UTC),
	}

	authorRepo.On("FindAuthorById", mock.Anything, db, int(authorID)).Return(Author, nil)
	service := NewAuthorService(authorRepo, db)

	result, err := service.FindDetailAuthor(context.Background(), int(authorID))

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, authorID, result.ID)
	assert.Equal(t, "Test Author", result.Name)
	assert.Equal(t, "1985-04-05", result.Birthdate)

	authorRepo.AssertExpectations(t)
}

func TestFindDetailAuthor_AuthorNotFound(t *testing.T) {
	authorRepo := new(repositories.MockAuthorRepository)
	db := new(gorm.DB)
	authorID := uint(1)

	authorRepo.On("FindAuthorById", mock.Anything, db, int(authorID)).Return(nil, errors.New("author not found"))
	service := NewAuthorService(authorRepo, db)

	result, err := service.FindDetailAuthor(context.Background(), int(authorID))

	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "NOT FOUND ERROR", err.Message)
	assert.Equal(t, 400, err.StatusCode)

	authorRepo.AssertExpectations(t)
}

func TestFindAllAuthor_Success(t *testing.T) {
	authorRepo := new(repositories.MockAuthorRepository)
	db := new(gorm.DB)

	Authors := []*models.Author{
		{
			ID:        1,
			Name:      "Test Author",
			Birthdate: time.Date(1985, time.April, 5, 0, 0, 0, 0, time.UTC),
		},
	}

	authorRepo.On("GetListAuthors", mock.Anything, db).Return(Authors, nil)
	service := NewAuthorService(authorRepo, db)

	result, err := service.FindAllAuthors(context.Background())

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, "Test Author", result[0].Name)
	assert.Equal(t, "1985-04-05", result[0].Birthdate)

	authorRepo.AssertExpectations(t)
}

func TestFindAllAuthors_RepositoryError(t *testing.T) {
	authorRepo := new(repositories.MockAuthorRepository)
	db := new(gorm.DB)
	service := NewAuthorService(authorRepo, db)

	authorRepo.On("GetListAuthors", mock.Anything, db).Return(nil, errors.New("db error"))

	result, err := service.FindAllAuthors(context.Background())

	assert.NotNil(t, err)
	assert.Nil(t, result)
	authorRepo.AssertExpectations(t)
}

func TestCreateAuthor_Success(t *testing.T) {
	authorRepo := new(repositories.MockAuthorRepository)
	db := new(gorm.DB)
	service := NewAuthorService(authorRepo, db)

	validRequest := &params.AuthorRequest{
		Name:      "Test Author",
		Birthdate: "2001-03-24",
	}

	authorRepo.On("CreateAuthor", mock.Anything, db, mock.AnythingOfType("*models.Author")).Return(nil)

	errCust := service.CrateAuthor(context.Background(), validRequest)

	assert.Nil(t, errCust)
	authorRepo.AssertExpectations(t)
}

func TestCreateAuthor_ValidationError(t *testing.T) {
	authorRepo := new(repositories.MockAuthorRepository)
	db := new(gorm.DB)
	service := NewAuthorService(authorRepo, db)

	invalidRequest := &params.AuthorRequest{
		Name:      "Test Author",
		Birthdate: "",
	}

	err := service.CrateAuthor(context.Background(), invalidRequest)

	assert.NotNil(t, err)
	assert.Equal(t, "BAD REQUEST ERROR", err.Message)
}

func TestCreateAuthor_RepositoryError(t *testing.T) {
	authorRepo := new(repositories.MockAuthorRepository)
	db := new(gorm.DB)
	service := NewAuthorService(authorRepo, db)

	validRequest := &params.AuthorRequest{
		Name:      "Test Author",
		Birthdate: "1985-04-05",
	}

	authorRepo.On("CreateAuthor", mock.Anything, db, mock.AnythingOfType("*models.Author")).Return(errors.New("db error"))

	err := service.CrateAuthor(context.Background(), validRequest)

	assert.NotNil(t, err)
	assert.Equal(t, "BAD REQUEST ERROR", err.Message)
	authorRepo.AssertExpectations(t)
}

func TestUpdateAuthor_Success(t *testing.T) {
	authorRepo := new(repositories.MockAuthorRepository)
	db := new(gorm.DB)
	service := NewAuthorService(authorRepo, db)

	validRequest := &params.AuthorRequest{
		Name:      "Update Author",
		Birthdate: "1985-04-05",
	}

	authorRepo.On("UpdateAuthor", mock.Anything, db, mock.AnythingOfType("*models.Author")).Return(nil)

	result, err := service.UpdateAuthor(context.Background(), 1, validRequest)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Update Author", result.Name)
	authorRepo.AssertExpectations(t)
}

func TestUpdateAuthor_ValidationError(t *testing.T) {
	authorRepo := new(repositories.MockAuthorRepository)
	db := new(gorm.DB)
	service := NewAuthorService(authorRepo, db)

	invalidRequest := &params.AuthorRequest{
		Name:      "",
		Birthdate: "1985-04-05",
	}

	result, err := service.UpdateAuthor(context.Background(), 1, invalidRequest)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "BAD REQUEST ERROR", err.Message)
}

func TestUpdateAuthor_RepositoryError(t *testing.T) {
	authorRepo := new(repositories.MockAuthorRepository)
	db := new(gorm.DB)
	service := NewAuthorService(authorRepo, db)

	validRequest := &params.AuthorRequest{
		Name:      "Update Author",
		Birthdate: "1985-04-05",
	}

	authorRepo.On("UpdateAuthor", mock.Anything, db, mock.AnythingOfType("*models.Author")).Return(errors.New("db error"))

	result, err := service.UpdateAuthor(context.Background(), 1, validRequest)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "BAD REQUEST ERROR", err.Message)
	authorRepo.AssertExpectations(t)
}

func TestDeleteAuthor_Success(t *testing.T) {
	authorRepo := new(repositories.MockAuthorRepository)
	db := new(gorm.DB)
	service := NewAuthorService(authorRepo, db)

	authorRepo.On("DeleteAuthor", mock.Anything, db, 1).Return(nil)

	err := service.DeleteAuthor(context.Background(), 1)

	assert.Nil(t, err)
	authorRepo.AssertExpectations(t)
}

func TestDeleteAuthor_RepositoryError(t *testing.T) {
	authorRepo := new(repositories.MockAuthorRepository)
	db := new(gorm.DB)
	service := NewAuthorService(authorRepo, db)

	authorRepo.On("DeleteAuthor", mock.Anything, db, 1).Return(errors.New("author not found"))

	err := service.DeleteAuthor(context.Background(), 1)

	assert.NotNil(t, err)
	assert.Equal(t, "NOT FOUND ERROR", err.Message)
	authorRepo.AssertExpectations(t)
}
