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
	books := model.Book{}
	books.Title = "Clean Code"
	books.Author = "Robert C.Martin"
	books.Page = 750
	books.Price = 20.1
	books.Page = 750
	books.Release = 1997

	er := db.Create(&books).Error
	if er != nil {
		log.Fatal("create failed")
	}

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
