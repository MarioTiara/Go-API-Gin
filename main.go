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

	// Get the first record ordered by primary key
	var book1 model.Book
	err = db.First(&book1).Error
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(book1)
	}
	// Get one record, no specified order
	var book2 model.Book
	err = db.Take(&book2).Error
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(book2)
	}

	// Get last record, ordered by primary key desc
	var book3 model.Book
	err = db.Last(&book3).Error
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(book3)
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
