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
		bookRouter.GET(":id/:code", controller.BookHanlderMultiParam)
	}

	router.Run(":888")
}
