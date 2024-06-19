package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

var secKey = []byte("your_secret_key")

func tokenGenerator(username string, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["role"] = role
	claims["exp_time"] = time.Now().Add(time.Hour * 24).Unix()

	return token.SignedString(secKey)
}

func authenticMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			ok := token.Method.(*jwt.SigningMethodHMAC)
			if ok != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secKey, nil
		})

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token is invalid"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		validToken := token.Valid

		if ok && validToken {
			c.Set("claims", claims)
		}

		c.Next()

	}
}

func roleBasedMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		userRole := claims.(jwt.MapClaims)["role"].(string)

		if requiredRole != userRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "you are not allowed to access this resource"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		username := "raju657"
		role := "admin"
		tokenString, _ := tokenGenerator(username, role)
		c.JSON(http.StatusOK, gin.H{"Your token": tokenString})
	})

	protected := router.Group("/protected")
	protected.Use(authenticMiddleware())
	{
		protected.GET("/admin", roleBasedMiddleware("admin"), func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Hello Admin!"})
		})

		protected.GET("/user", roleBasedMiddleware("user"), func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Hello User!"})
		})
	}

	router.Run(":8080")
}
