package controllers

import (
	"golang-backend-test/app/commons/response"
	"golang-backend-test/app/params"
	"golang-backend-test/app/services"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Register(ginCtx *gin.Context)
	Login(ginCtx *gin.Context)
}

type UserControllerImpl struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Register(ginCtx *gin.Context) {
	var request = new(params.UserRequest)
	err := ginCtx.ShouldBindJSON(request)
	if err != nil {
		errParam := response.GeneralError()
		ginCtx.AbortWithStatusJSON(errParam.StatusCode, errParam)
		return
	}

	custErr := controller.UserService.Register(ginCtx, request)
	if custErr != nil {
		ginCtx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}
	resp := response.CreatedSuccessCustomMessageAndPayload("Success register users", nil)
	ginCtx.JSON(resp.StatusCode, resp)
}

func (controller *UserControllerImpl) Login(ginCtx *gin.Context) {
	var request = new(params.UserRequest)
	err := ginCtx.ShouldBindJSON(request)
	if err != nil {
		errParam := response.GeneralError()
		ginCtx.AbortWithStatusJSON(errParam.StatusCode, errParam)
		return
	}

	result, custErr := controller.UserService.Login(ginCtx, request)
	if custErr != nil {
		ginCtx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}
	resp := response.GeneralSuccessCustomMessageAndPayload("Success login users", result)
	ginCtx.JSON(resp.StatusCode, resp)
}
