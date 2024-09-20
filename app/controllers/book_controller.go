package controllers

import (
	"golang-backend-test/app/commons/response"
	"golang-backend-test/app/params"
	"golang-backend-test/app/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController interface {
	FindBookById(ginCtx *gin.Context)
	GetListBooks(ginCtx *gin.Context)
	CreateBook(ginCtx *gin.Context)
	UpdateBook(ginCtx *gin.Context)
	DeleteBook(ginCtx *gin.Context)
}

type BookControllerImpl struct {
	BookService services.BookService
}

func NewBookController(bookService services.BookService) BookController {
	return &BookControllerImpl{
		BookService: bookService,
	}
}

func (controller *BookControllerImpl) FindBookById(ginCtx *gin.Context) {
	id, err := strconv.Atoi(ginCtx.Param("id"))
	if err != nil {
		errParam := response.NotFoundError()
		ginCtx.AbortWithStatusJSON(errParam.StatusCode, errParam)
		return
	}
	result, custErr := controller.BookService.FindDetailBook(ginCtx, id)
	if custErr != nil {
		ginCtx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.GeneralSuccessCustomMessageAndPayload("Success get data detail books.", result)
	ginCtx.JSON(resp.StatusCode, resp)
}

func (controller *BookControllerImpl) GetListBooks(ginCtx *gin.Context) {
	result, custErr := controller.BookService.FindAllBooks(ginCtx)
	if custErr != nil {
		ginCtx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.GeneralSuccessCustomMessageAndPayload("Success get data books.", result)
	ginCtx.JSON(resp.StatusCode, resp)
}

func (controller *BookControllerImpl) CreateBook(ginCtx *gin.Context) {
	var request = new(params.BookRequest)
	err := ginCtx.ShouldBindJSON(request)
	if err != nil {
		errParam := response.GeneralError()
		ginCtx.AbortWithStatusJSON(errParam.StatusCode, errParam)
		return
	}

	custErr := controller.BookService.CrateBook(ginCtx, request)
	if custErr != nil {
		ginCtx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}
	resp := response.CreatedSuccess()
	ginCtx.JSON(resp.StatusCode, resp)
}

func (controller *BookControllerImpl) UpdateBook(ginCtx *gin.Context) {
	var request = new(params.BookRequest)
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

	result, custErr := controller.BookService.UpdateBook(ginCtx, id, request)
	if custErr != nil {
		ginCtx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}
	resp := response.GeneralSuccessCustomMessageAndPayload("Success update data books", result)
	ginCtx.JSON(resp.StatusCode, resp)
}

func (controller *BookControllerImpl) DeleteBook(ginCtx *gin.Context) {
	id, err := strconv.Atoi(ginCtx.Param("id"))
	if err != nil {
		errParam := response.NotFoundError()
		ginCtx.AbortWithStatusJSON(errParam.StatusCode, errParam)
		return
	}

	custErr := controller.BookService.DeleteBook(ginCtx, id)
	if custErr != nil {
		ginCtx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}
	resp := response.GeneralSuccess()
	ginCtx.JSON(resp.StatusCode, resp)
}
