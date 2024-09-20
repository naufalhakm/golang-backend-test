package controllers

import (
	"golang-backend-test/app/commons/response"
	"golang-backend-test/app/params"
	"golang-backend-test/app/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthorController interface {
	FindAuthorById(ginCtx *gin.Context)
	GetListAuthors(ginCtx *gin.Context)
	CreateAuthor(ginCtx *gin.Context)
	UpdateAuthor(ginCtx *gin.Context)
	DeleteAuthor(ginCtx *gin.Context)
}

type AuthorControllerImpl struct {
	AuthorService services.AuthorService
}

func NewAuthorController(authorService services.AuthorService) AuthorController {
	return &AuthorControllerImpl{
		AuthorService: authorService,
	}
}

func (controller *AuthorControllerImpl) FindAuthorById(ginCtx *gin.Context) {
	id, err := strconv.Atoi(ginCtx.Param("id"))
	if err != nil {
		errParam := response.NotFoundError()
		ginCtx.AbortWithStatusJSON(errParam.StatusCode, errParam)
		return
	}
	result, custErr := controller.AuthorService.FindDetailAuthor(ginCtx, id)
	if custErr != nil {
		ginCtx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.GeneralSuccessCustomMessageAndPayload("Success get data detail authors.", result)
	ginCtx.JSON(resp.StatusCode, resp)
}

func (controller *AuthorControllerImpl) GetListAuthors(ginCtx *gin.Context) {
	result, custErr := controller.AuthorService.FindAllAuthors(ginCtx)
	if custErr != nil {
		ginCtx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.GeneralSuccessCustomMessageAndPayload("Success get data authors.", result)
	ginCtx.JSON(resp.StatusCode, resp)
}

func (controller *AuthorControllerImpl) CreateAuthor(ginCtx *gin.Context) {
	var request = new(params.AuthorRequest)
	err := ginCtx.ShouldBindJSON(request)
	if err != nil {
		errParam := response.GeneralError()
		ginCtx.AbortWithStatusJSON(errParam.StatusCode, errParam)
		return
	}

	custErr := controller.AuthorService.CrateAuthor(ginCtx, request)
	if custErr != nil {
		ginCtx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}
	resp := response.CreatedSuccess()
	ginCtx.JSON(resp.StatusCode, resp)
}

func (controller *AuthorControllerImpl) UpdateAuthor(ginCtx *gin.Context) {
	var request = new(params.AuthorRequest)
	err := ginCtx.ShouldBindJSON(request)
	if err != nil {
		errParam := response.GeneralError()
		ginCtx.AbortWithStatusJSON(errParam.StatusCode, errParam)
		return
	}
	id, err := strconv.Atoi(ginCtx.Param("id"))
	if err != nil {
		errParam := response.NotFoundError()
		ginCtx.AbortWithStatusJSON(errParam.StatusCode, errParam)
		return
	}

	result, custErr := controller.AuthorService.UpdateAuthor(ginCtx, id, request)
	if custErr != nil {
		ginCtx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}
	resp := response.GeneralSuccessCustomMessageAndPayload("Success update data authors", result)
	ginCtx.JSON(resp.StatusCode, resp)
}

func (controller *AuthorControllerImpl) DeleteAuthor(ginCtx *gin.Context) {
	id, err := strconv.Atoi(ginCtx.Param("id"))
	if err != nil {
		errParam := response.NotFoundError()
		ginCtx.AbortWithStatusJSON(errParam.StatusCode, errParam)
		return
	}

	custErr := controller.AuthorService.DeleteAuthor(ginCtx, id)
	if custErr != nil {
		ginCtx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}
	resp := response.GeneralSuccess()
	ginCtx.JSON(resp.StatusCode, resp)
}
