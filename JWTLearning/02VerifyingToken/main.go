package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("your_secret_key")

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	// Set claims (payload)
	claims["user_id"] = "123"
	claims["username"] = "john_doe"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Generate encoded token and return
	return token.SignedString(secretKey)
}

func verifyJWT(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key
		return secretKey, nil
	})
	if err != nil {
		return fmt.Errorf("error parsing token: %v", err)
	}

	if !token.Valid {
		return fmt.Errorf("token is invalid")
	}

	// Token is valid
	return nil
}

func main() {
	tokenString, err := generateJWT()
	if err != nil {
		fmt.Println("Error generating JWT:", err)
		return
	}
	fmt.Println("Generated JWT:", tokenString)

	// Verify the JWT
	err = verifyJWT(tokenString)
	if err != nil {
		fmt.Println("Error verifying JWT:", err)
		return
	}
	fmt.Println("JWT is valid")
}
