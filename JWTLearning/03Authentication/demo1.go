package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

var secKey = []byte("personal_secret_key")

func authentication(c *gin.Context) {
	token, _ := tokenGenerator(1, "hariom")
	c.IndentedJSON(http.StatusOK, gin.H{"Your token": token})
}

func tokenGenerator(id int, username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = id
	claims["username"] = username
	claims["exp_time"] = time.Now().Add(time.Hour * 24).Unix()

	return token.SignedString(secKey)
}

func authMiddlewarePro() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "He/She has no permission or taken to access!"})
			c.Abort()
			return
		}

		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secKey, nil
		})

		if !token.Valid {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"Error": "He/She has no permission or token invalid!"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func protectedCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello! Your are authorized---"})
}
func main() {
	router := gin.Default()
	router.POST("/login", authentication)

	protected := router.Group("/protected")
	protected.Use(authMiddlewarePro())
	protected.GET("/", protectedCheck)
	router.Run(":8080")
}
