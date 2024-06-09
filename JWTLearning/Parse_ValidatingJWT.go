package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func main() {
	// Your token string
	tokenString := "your.jwt.token.here"

	// Keyfunc to provide the key
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		// Ensure the token's signing method matches our expectation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("your-256-bit-secret"), nil
	}

	// Create a new parser with options
	parser := jwt.NewParser(
		jwt.WithValidMethods([]string{"HS256"}),
		jwt.WithIssuer("my-auth-server"),
		jwt.WithAudience("my-audience"),
		jwt.WithLeeway(5*time.Minute),
		jwt.WithIssuedAt(),
	)

	// Parse and validate the token
	token, err := parser.Parse(tokenString, keyFunc)
	if err != nil {
		fmt.Println("Error parsing token:", err)
		return
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("Token is valid!")
		fmt.Println("Claims:", claims)
	} else {
		fmt.Println("Invalid token")
	}
}
