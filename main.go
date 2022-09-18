package main

import (
	"github.com/MarioTiara/Go-API-Gin/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", controller.HomeHanlder)
	router.Run(":888")
}
