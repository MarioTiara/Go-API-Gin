package authentication

import "github.com/golang-jwt/jwt/v4"

type Crediantials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var JwtKey = []byte("secret_key")
