package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

var secretKey = []byte("my_secret_key")

func generateJWT(username string, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = 7937
	claims["username"] = username
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	return token.SignedString(secretKey)
}

func main() {
	tokenString, err := generateJWT("Hariom Shah", "admin")
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated JWT:", tokenString)
}
