package services

import (
	"context"
	"golang-backend-test/app/commons/response"
	"golang-backend-test/app/models"
	"golang-backend-test/app/params"
	"golang-backend-test/app/repositories"
	"golang-backend-test/pkg/encryption"
	"golang-backend-test/pkg/token"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type UserService interface {
	Register(ctx context.Context, req *params.UserRequest) *response.CustomError
	Login(ctx context.Context, req *params.UserRequest) (*params.UserResponse, *response.CustomError)
}

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
	DB             *gorm.DB
}

func NewUserService(userRepository repositories.UserRepository, db *gorm.DB) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
	}
}

func (service *UserServiceImpl) Register(ctx context.Context, req *params.UserRequest) *response.CustomError {
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

	_, errCheck := service.UserRepository.FindUserByUsername(ctx, service.DB, req.Username)
	if errCheck == nil {
		return response.BadRequestErrorWithAdditionalInfo("username already exists")
	}

	hashPaswword, err := encryption.HashPassword(req.Password)
	if err != nil {
		return response.GeneralError()
	}

	var user = new(models.User)
	user.Username = req.Username
	user.Password = hashPaswword
	if err := service.UserRepository.CreateUser(ctx, service.DB, user); err != nil {
		return response.BadRequestError()
	}

	return nil
}

func (service *UserServiceImpl) Login(ctx context.Context, req *params.UserRequest) (*params.UserResponse, *response.CustomError) {
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

	user, err := service.UserRepository.FindUserByUsername(ctx, service.DB, req.Username)
	if err != nil {
		return nil, response.NotFoundError()
	}

	verif := encryption.VerifyPassword(req.Password, user.Password)
	if !verif {

		return nil, response.GeneralError()
	}
	token, err := token.GenerateToken(int(user.ID))
	if err != nil {
		return nil, response.GeneralErrorWithAdditionalInfo(err.Error())
	}

	return &params.UserResponse{
		Token: token,
	}, nil
}
