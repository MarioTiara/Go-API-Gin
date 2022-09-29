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

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (handler *bookHandler) BookHandler(c *gin.Context) {
	query := c.Query("id")
	if len(query) > 0 {
		id, err := strconv.Atoi(query)
		if err != nil {
			log.Fatal(err)
		} else {
			book, err := handler.bookService.FindByID(id)
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
		books, _ := handler.bookService.FindAll()
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"books":  books,
		})
	}

}

func (handler *bookHandler) BookhandlerUrlParam(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	} else {
		book, err := handler.bookService.FindByID(id)
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

func (handler *bookHandler) PostBookHadler(c *gin.Context) {
	var newbook book.BookRequest
	err := c.ShouldBindJSON(&newbook)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("ërror on field %s, condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errorMessage)
			return
		}
	} else {
		book, err := handler.bookService.Create(newbook)
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

func (handler *bookHandler) DeleteBookhandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	} else {
		err := handler.bookService.DeleteByID(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			c.JSON(http.StatusOK, nil)
		}
	}
}

func (handler *bookHandler) UpdateHandler(c *gin.Context) {
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
		err := handler.bookService.Update(incomingBook)
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
