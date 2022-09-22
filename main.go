package main

import (
	"fmt"
	"log"

	"github.com/MarioTiara/Go-API-Gin/controller"
	"github.com/MarioTiara/Go-API-Gin/model"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=Mario2022! dbname=Product port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		message := fmt.Sprintf("Db connection error: %s", err.Error())
		log.Fatal(message)
	}

	fmt.Println("Database Connection succeed")
	db.AutoMigrate(&model.Book{})
	db.Delete(&model.Book{}, 4)

	router := gin.Default()
	router.GET("/", controller.HomeHandler)
	bookRouter := router.Group("book")
	{
		bookRouter.GET("", controller.BooksHanlder)
		bookRouter.GET(":id", controller.BookHanlder)
		bookRouter.GET(":id/:code", controller.BookHandlerMultiParam)
		bookRouter.POST("", controller.PostBookHadler)
	}

	router.Run(":888")
}
