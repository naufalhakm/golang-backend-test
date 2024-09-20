package services

import (
	"context"
	"errors"
	"fmt"
	"golang-backend-test/app/models"
	"golang-backend-test/app/params"
	"golang-backend-test/app/repositories"
	"golang-backend-test/pkg/encryption"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestRegister_Success(t *testing.T) {
	// Arrange
	mockRepo := new(repositories.MockUserRepository)
	db := new(gorm.DB)
	service := NewUserService(mockRepo, db)

	validRequest := &params.UserRequest{
		Username: "naufalhakm",
		Password: "password123",
	}

	mockRepo.On("FindUserByUsername", mock.Anything, db, "naufalhakm").Return(nil, errors.New("users not found"))

	mockRepo.On("CreateUser", mock.Anything, db, mock.AnythingOfType("*models.User")).Return(nil)

	err := service.Register(context.Background(), validRequest)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRegister_ValidationErrorRequired(t *testing.T) {
	// Arrange
	mockRepo := new(repositories.MockUserRepository)
	db := new(gorm.DB)
	service := NewUserService(mockRepo, db)

	invalidRequest := &params.UserRequest{
		Username: "",
		Password: "password",
	}

	err := service.Register(context.Background(), invalidRequest)

	assert.NotNil(t, err)
	assert.Equal(t, 400, err.StatusCode)
	assert.Equal(t, "BAD REQUEST ERROR", err.Message)
}

func TestRegister_ValidationErrorMaximal(t *testing.T) {
	// Arrange
	mockRepo := new(repositories.MockUserRepository)
	db := new(gorm.DB)
	service := NewUserService(mockRepo, db)

	invalidRequest := &params.UserRequest{
		Username: "naufalhakm",
		Password: "qwerty",
	}

	err := service.Register(context.Background(), invalidRequest)

	assert.NotNil(t, err)
	assert.Equal(t, 400, err.StatusCode)
	assert.Equal(t, "BAD REQUEST ERROR", err.Message)
}

func TestRegister_UserAlreadyExists(t *testing.T) {
	// Arrange
	mockRepo := new(repositories.MockUserRepository)
	db := new(gorm.DB)
	service := NewUserService(mockRepo, db)

	invalidRequest := &params.UserRequest{
		Username: "naufalhakm",
		Password: "qwerty",
	}

	user := &models.User{
		ID:       1,
		Username: "naufalhakm",
		Password: "hashedpassword",
	}

	mockRepo.On("FindUserByUsername", mock.Anything, db, "naufalhakm").Return(user, nil)

	err := service.Register(context.Background(), invalidRequest)

	assert.NotNil(t, err)
	assert.Equal(t, 400, err.StatusCode)
	assert.Equal(t, "BAD REQUEST ERROR", err.Message)

}

func TestLogin_Success(t *testing.T) {
	mockRepo := new(repositories.MockUserRepository)
	db := new(gorm.DB)
	service := NewUserService(mockRepo, db)

	validRequest := &params.UserRequest{
		Username: "naufalhakm",
		Password: "password123",
	}

	hashPaswword, _ := encryption.HashPassword(validRequest.Password)

	user := &models.User{
		ID:       1,
		Username: "naufalhakm",
		Password: hashPaswword,
	}

	mockRepo.On("FindUserByUsername", mock.Anything, db, "naufalhakm").Return(user, nil)

	result, err := service.Login(context.Background(), validRequest)

	fmt.Println(err)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, result.Token)

	mockRepo.AssertExpectations(t)
}

func TestLogin_UserNotFound(t *testing.T) {
	// Arrange
	mockRepo := new(repositories.MockUserRepository)
	db := new(gorm.DB)
	service := NewUserService(mockRepo, db)

	invalidRequest := &params.UserRequest{
		Username: "invaliduser",
		Password: "password",
	}

	mockRepo.On("FindUserByUsername", mock.Anything, db, "invaliduser").Return(nil, errors.New("user not found"))

	_, err := service.Login(context.Background(), invalidRequest)

	assert.NotNil(t, err)
	assert.Equal(t, "NOT FOUND ERROR", err.Message)
	assert.Equal(t, 400, err.StatusCode)
	mockRepo.AssertExpectations(t)
}

func TestLogin_ValidationErrorRequired(t *testing.T) {
	// Arrange
	mockRepo := new(repositories.MockUserRepository)
	db := new(gorm.DB)
	service := NewUserService(mockRepo, db)

	invalidRequest := &params.UserRequest{
		Username: "",
		Password: "password",
	}

	_, err := service.Login(context.Background(), invalidRequest)

	assert.NotNil(t, err)
	assert.Equal(t, "BAD REQUEST ERROR", err.Message)
	assert.Equal(t, 400, err.StatusCode)
}

func TestLogin_ValidationErrorMinimun(t *testing.T) {
	// Arrange
	mockRepo := new(repositories.MockUserRepository)
	db := new(gorm.DB)
	service := NewUserService(mockRepo, db)

	invalidRequest := &params.UserRequest{
		Username: "naufalhakm",
		Password: "naufal",
	}

	_, err := service.Login(context.Background(), invalidRequest)

	assert.NotNil(t, err)
	assert.Equal(t, "BAD REQUEST ERROR", err.Message)
	assert.Equal(t, 400, err.StatusCode)
}
