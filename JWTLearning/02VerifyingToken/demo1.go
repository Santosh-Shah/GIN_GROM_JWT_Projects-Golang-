package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

var secKey = []byte("this_is_secret_key")

func tokenGenerator() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = 425
	claims["use_name"] = "raju456"
	claims["exp_time"] = time.Now().Add(time.Hour * 24).Unix()

	return token.SignedString(secKey)
}

func verificationToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secKey, nil
	})

	if err != nil {
		return fmt.Errorf("parse token: %v", err)
	}

	if !token.Valid {
		return fmt.Errorf("token is invalid")
	}
	return nil
}

func main() {
	tokenValue, _ := tokenGenerator()
	fmt.Println("tokenValue:", tokenValue)

	err := verificationToken(tokenValue)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("verificationToken Ok/Valid")
	}
}
