package main

import (
	"github.com/MarioTiara/Go-API-Gin/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", controller.HomeHandler)
	router.GET("books", controller.BooksHanlder)
	router.Run(":888")
}
