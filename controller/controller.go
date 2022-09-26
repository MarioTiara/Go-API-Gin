package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/MarioTiara/Go-API-Gin/data"
	"github.com/MarioTiara/Go-API-Gin/model"
	"github.com/MarioTiara/Go-API-Gin/postgres"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var db, err = postgres.GetDb()
var repository = data.NewRepository(db)

func HomeHandler(c *gin.Context) {
	token := c.Request.Header["Token"]
	c.JSON(http.StatusOK, gin.H{
		"status":     "OK",
		"message":    "Home page",
		"token_data": token,
	})
}

func BooksHanlder(c *gin.Context) {
	query := c.Query("id")
	if len(query) > 0 {
		id, err := strconv.Atoi(query)
		if err != nil {
			log.Fatal(err)
		} else {
			book, err := repository.FindByID(id)
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
		books, _ := repository.FindAll()
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"books":  books,
		})
	}

}

func BookHanlder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	} else {
		book, err := repository.FindByID(id)
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
	var newbook model.Book
	err := c.ShouldBindJSON(&newbook)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("ërror on field %s, condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errorMessage)
			return
		}
	} else {
		book, err := repository.Create(newbook)
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

func DeleteBookHanlder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	} else {
		err := repository.DeleteById(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			c.JSON(http.StatusOK, nil)
		}
	}
}

func UpdateHandler(c *gin.Context) {
	var incomingBook model.Book
	err := c.ShouldBindJSON(&incomingBook)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("errr on field %s, condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errMessage)
			return
		}
	} else {
		err := repository.Update(incomingBook)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			message := fmt.Sprintf("Book with Id:%d has been update", incomingBook.Id)
			c.JSON(http.StatusOK, gin.H{
				"status":  "OK",
				"message": message,
			})
		}
	}
}
