package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/MarioTiara/Go-API-Gin/data"
	"github.com/MarioTiara/Go-API-Gin/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HomeHandler(c *gin.Context) {
	token := c.Request.Header["Token"]
	c.JSON(http.StatusOK, gin.H{
		"status":     "OK",
		"message":    "Home page",
		"token_data": token,
	})
}

func BooksHanlder(c *gin.Context) {
	data := &data.DbBooks
	query := c.Query("id")
	if len(query) > 0 {
		id, err := strconv.Atoi(query)
		if err != nil {
			log.Fatal(err)
		} else {
			if id < len(data.Item) {
				c.JSON(http.StatusOK, gin.H{
					"status": "OK",
					"books":  data.Item[id],
				})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "Bad Request",
					"message": "Your Id is out of index",
				})
			}
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"books":  data,
		})
	}

}

func BookHanlder(c *gin.Context) {
	data := &data.DbBooks
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	if len(data.Item) > id {
		book := data.Item[id]
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"book":   book,
		})
	}
}

func BookHandlerMultiParam(c *gin.Context) {
	data := &data.DbBooks
	id, _ := strconv.Atoi(c.Param("id"))
	code := c.Param("code")
	if id < len(data.Item) && len(code) > 0 {
		book := data.Item[id]
		if book.Code == code {
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
				"book":   book,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "Bad Request",
				"message": fmt.Sprintf("There is no book with id: %d and code: %s", id, code),
			})
		}
	}
}

func PostBookHadler(c *gin.Context) {
	var newbook *model.Book
	dbBooks := &data.DbBooks
	err := c.ShouldBindJSON(&newbook)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("ërror on field %s, condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errorMessage)
			return
		}
	}
	if err == nil {
		dbBooks.Item = append(dbBooks.Item, *newbook)
		c.IndentedJSON(http.StatusOK, dbBooks.Item[len(dbBooks.Item)-1])

	}

}
