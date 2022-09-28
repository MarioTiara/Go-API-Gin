package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/MarioTiara/Go-API-Gin/book"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// var db, _ = postgres.GetDb()
// var bookRepository = book.NewRepository(db)
// var bookService = book.NewService(bookRepository)

type bookHanlder struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHanlder {
	return &bookHanlder{bookService}
}

func (hanlder *bookHanlder) BooksHanlder(c *gin.Context) {
	query := c.Query("id")
	if len(query) > 0 {
		id, err := strconv.Atoi(query)
		if err != nil {
			log.Fatal(err)
		} else {
			book, err := hanlder.bookService.FindByID(id)
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
		books, _ := hanlder.bookService.FindAll()
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"books":  books,
		})
	}

}

func (hanlder *bookHanlder) BookHanlderUrlParam(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	} else {
		book, err := hanlder.bookService.FindByID(id)
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

func (hanlder *bookHanlder) PostBookHadler(c *gin.Context) {
	var newbook book.BookRequest
	err := c.ShouldBindJSON(&newbook)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("ërror on field %s, condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errorMessage)
			return
		}
	} else {
		book, err := hanlder.bookService.Create(newbook)
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

func (hanlder *bookHanlder) DeleteBookHanlder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	} else {
		err := hanlder.bookService.DeleteByID(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			c.JSON(http.StatusOK, nil)
		}
	}
}

func (hanlder *bookHanlder) UpdateHandler(c *gin.Context) {
	var incomingBook book.BookRequest
	err := c.ShouldBindJSON(&incomingBook)
	id, _ := incomingBook.ID.Int64()
	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "the ID can't 0",
		})
	}

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("errr on field %s, condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errMessage)
			return
		}
	} else {
		err := hanlder.bookService.Update(incomingBook)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			message := fmt.Sprintf("Book with Title:%s has been update", incomingBook.Title)
			c.JSON(http.StatusOK, gin.H{
				"status":  "OK",
				"message": message,
			})
		}
	}
}
