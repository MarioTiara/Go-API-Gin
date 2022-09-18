package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/MarioTiara/Go-API-Gin/data"
	"github.com/gin-gonic/gin"
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

func BookHanlderMultiParam(c *gin.Context) {
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
