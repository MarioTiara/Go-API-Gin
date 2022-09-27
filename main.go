package main

import (
	"net/http"

	"github.com/MarioTiara/Go-API-Gin/book"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", HomeHandler)
	bookRouter := router.Group("book")
	{
		bookRouter.GET("", book.BooksHanlder)
		bookRouter.GET(":id", book.BookHanlderUrlParam)
		bookRouter.POST("", book.PostBookHadler)
	}

	router.Run(":888")
}

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"massage": "welcome Gin API Project by Mario",
	})
}
