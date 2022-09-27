package book

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/MarioTiara/Go-API-Gin/postgres"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var db, _ = postgres.GetDb()
var bookRepository = NewRepository(db)
var bookService = NewService(bookRepository)

func BooksHanlder(c *gin.Context) {
	query := c.Query("id")
	if len(query) > 0 {
		id, err := strconv.Atoi(query)
		if err != nil {
			log.Fatal(err)
		} else {
			book, err := bookService.FindByID(id)
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status": "OK",
					"books":  book,
				})
			}
		}
	} else {
		books, _ := bookService.FinAll()
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"books":  books,
		})
	}

}

func BookHanlderUrlParam(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	} else {
		book, err := bookService.FindByID(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
				"book":   book,
			})
		}
	}
}

func PostBookHadler(c *gin.Context) {
	var newbook Book
	err := c.ShouldBindJSON(&newbook)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("ërror on field %s, condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errorMessage)
			return
		}
	} else {
		book, err := bookRepository.Create(newbook)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"book":   book,
				"status": "OK",
			})
		}
	}

}
