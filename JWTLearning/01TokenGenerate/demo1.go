package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

var secKey = []byte("my_secret_key")

func tokenGenerator() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = 4526
	claims["user_name"] = "hariom789"
	claims["exp_time"] = time.Now().Add(time.Hour * 24).Unix()

	return token.SignedString(secKey)
}

func main() {
	tokenString, _ := tokenGenerator()
	fmt.Println("tokenString:", tokenString)
}
