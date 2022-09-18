package controller

import (
	"net/http"

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
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"books":  data,
	})

}
