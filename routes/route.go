package routes

import (
	"golang-backend-test/app/commons/response"
	"golang-backend-test/factory"
	"golang-backend-test/pkg/token"
	"strings"

	"github.com/gin-gonic/gin"
)

func NewRoutes(router *gin.Engine, provider *factory.Provider) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", provider.UserProvider.Register)
		auth.POST("/login", provider.UserProvider.Login)
	}

	authors := router.Group("/authors", CheckAuth())
	{
		authors.GET("/", provider.AuthorProvider.GetListAuthors)
		authors.POST("/", provider.AuthorProvider.CreateAuthor)
		authors.GET("/:id", provider.AuthorProvider.FindAuthorById)
		authors.PUT("/:id", provider.AuthorProvider.UpdateAuthor)
		authors.DELETE("/:id", provider.AuthorProvider.DeleteAuthor)
	}

	books := router.Group("/books", CheckAuth())
	{
		books.GET("/", provider.BookProvider.GetListBooks)
		books.POST("/", provider.BookProvider.CreateBook)
		books.GET("/:id", provider.BookProvider.FindBookById)
		books.PUT("/:id", provider.BookProvider.UpdateBook)
		books.DELETE("/:id", provider.BookProvider.DeleteBook)
	}
}

func CheckAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")

		bearerToken := strings.Split(header, "Bearer ")

		if len(bearerToken) != 2 {
			resp := response.UnauthorizedErrorWithAdditionalInfo("len token must be 2")
			ctx.AbortWithStatusJSON(resp.StatusCode, resp)
			return
		}

		payload, err := token.ValidateToken(bearerToken[1])
		if err != nil {
			resp := response.UnauthorizedErrorWithAdditionalInfo(err.Error())
			ctx.AbortWithStatusJSON(resp.StatusCode, resp)
			return
		}
		ctx.Set("authId", payload.AuthId)
		ctx.Next()
	}
}
