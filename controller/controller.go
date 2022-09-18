package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHanlder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Home page",
	})
}
