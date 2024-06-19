package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

var secretKey = []byte("my_secret_key")

func generateJWT(userID int, username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	return token.SignedString(secretKey)
}

func authenticate(c *gin.Context) {
	// Normally you would validate the user credentials from a database
	// Here, for simplicity, we assume the user is authenticated
	tokenString, err := generateJWT(1, "sandip")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Request does not contain an access token"})
			c.Abort()
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secretKey, nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("error parsing token: %v", err)})
			c.Abort()
			return
		}
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token is invalid"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func protectedEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the protected endpoint!"})
}

func main() {
	router := gin.Default()
	router.POST("/login", authenticate)
	protected := router.Group("/protected")
	protected.Use(authMiddleware())
	protected.GET("/", protectedEndpoint)
	router.Run(":8080")
}
