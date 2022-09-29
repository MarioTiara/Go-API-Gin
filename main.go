package main

import (
	"net/http"

	"github.com/MarioTiara/Go-API-Gin/book"
	"github.com/MarioTiara/Go-API-Gin/handler"
	"github.com/MarioTiara/Go-API-Gin/postgres"
	"github.com/MarioTiara/Go-API-Gin/user"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", HomeHandler)
	db, _ := postgres.GetDb()
	db.AutoMigrate(&book.Book{})
	db.AutoMigrate(&user.User{})
	bookrepository := book.NewRepository(db)
	bookService := book.NewService(bookrepository)
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	bookHandler := handler.NewBookHandler(bookService)
	loginHandler := handler.NewLoginHandler(userService)

	router.POST("login", loginHandler.LoginHandler)
	bookRouter := router.Group("book")
	{
		bookRouter.GET("", bookHandler.BookHandler)
		bookRouter.GET(":id", bookHandler.BookhandlerUrlParam)
		bookRouter.POST("", bookHandler.PostBookHadler)
		bookRouter.DELETE(":id", bookHandler.DeleteBookhandler)
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
