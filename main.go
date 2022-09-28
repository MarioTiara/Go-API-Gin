package main

import (
	"net/http"

	"github.com/MarioTiara/Go-API-Gin/book"
	"github.com/MarioTiara/Go-API-Gin/handler"
	"github.com/MarioTiara/Go-API-Gin/postgres"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", HomeHandler)
	db, _ := postgres.GetDb()
	bookrepository := book.NewRepository(db)
	bookService := book.NewService(bookrepository)
	bookHandler := handler.NewBookHandler(bookService)

	bookRouter := router.Group("book")
	{
		bookRouter.GET("", bookHandler.BooksHanlder)
		bookRouter.GET(":id", bookHandler.BookHanlderUrlParam)
		bookRouter.POST("", bookHandler.PostBookHadler)
		bookRouter.DELETE(":id", bookHandler.DeleteBookHanlder)
		bookRouter.PUT("", bookHandler.UpdateHandler)
	}
	router.Run()
}

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"massage": "welcome Gin API Project by Mario",
	})
}
