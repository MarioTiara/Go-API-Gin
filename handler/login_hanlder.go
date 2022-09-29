package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"time"

	"github.com/MarioTiara/Go-API-Gin/authentication"
	"github.com/MarioTiara/Go-API-Gin/user"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// var users = map[string]string{
// 	"user1": "password1",
// 	"user2": "password2",
// }

// type Crediantials struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// type Claims struct {
// 	Username string `json:"username"`
// 	jwt.StandardClaims
// }

type loginHandler struct {
	userService user.Service
}

func NewLoginHandler(userService user.Service) *loginHandler {
	return &loginHandler{userService}
}

func (handler *loginHandler) LoginHandler(c *gin.Context) {
	var crediantials authentication.Crediantials
	err := json.NewDecoder(c.Request.Body).Decode(&crediantials)
	fmt.Println(crediantials)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := handler.userService.FindByName(crediantials.Username)
	if err != nil {
		c.Writer.WriteString("username is not found")
		return
	}

	if user.Password != crediantials.Password {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 5)
	claims := &authentication.Claims{
		Username: crediantials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(authentication.JwtKey)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"token":  tokenString,
	})

}
