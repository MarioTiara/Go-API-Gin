package main

import (
	"github.com/MarioTiara/Go-API-Gin/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", controller.HomeHandler)
	bookRouter := router.Group("book")
	{
		bookRouter.GET("", controller.BooksHanlder)
		bookRouter.GET(":id", controller.BookHanlder)
		bookRouter.POST("", controller.PostBookHadler)
		bookRouter.DELETE(":id", controller.DeleteBookHanlder)
	}

	router.Run(":888")
}
