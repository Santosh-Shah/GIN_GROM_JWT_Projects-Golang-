package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

//var secretKey = []byte("your_secret_key")

var secretKey = []byte("my_secret_key")

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = 7937
	claims["username"] = "Hariom Shah"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	return token.SignedString(secretKey)
}

func main() {
	tokenString, err := generateJWT()
	if err != nil {
		panic(err)
		return
	}

	fmt.Println("Generated JWT:", tokenString)
}

//func generateJWT() (string, error) {
//	token := jwt.New(jwt.SigningMethodHS256)
//	claims := token.Claims.(jwt.MapClaims)
//
//	// Set claims (payload)
//	claims["user_id"] = "4283"
//	claims["username"] = "Santosh Shah"
//	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours
//
//	// Generate encoded token and return
//	return token.SignedString(secretKey)
//}
//
//func main() {
//	tokenString, err := generateJWT()
//	if err != nil {
//		fmt.Println("Error generating JWT:", err)
//		return
//	}
//	fmt.Println("Generated JWT:", tokenString)
//}
