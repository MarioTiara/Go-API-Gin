package controller

import (
	"fmt"
	"net/http"

	"github.com/MarioTiara/Go-API-Gin/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func PostUserHandler(c *gin.Context) {
	var newuser model.User
	err := c.ShouldBindJSON(&newuser)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("ërror on field %s, condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errorMessage)
			return
		}
	} else {
		user, err := repository.Create(newuser)
	}
}
