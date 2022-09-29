package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("secret_key")

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Crediantials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func LoginHandler(c *gin.Context) {
	var crediantials Crediantials

	err := json.NewDecoder(c.Request.Body).Decode(&crediantials)
	fmt.Println(crediantials)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[crediantials.Username]
	if !ok || expectedPassword != crediantials.Password {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &Claims{
		Username: crediantials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"token":  tokenString,
	})

}
