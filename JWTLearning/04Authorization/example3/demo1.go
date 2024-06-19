package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

var secKey = []byte("your_secret_key")

func tokenGenerator(username string, userrole string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["role"] = userrole
	claims["exp_time"] = time.Now().Add(time.Hour * 24).Unix()
	return token.SignedString(secKey)
}

func authHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secKey, nil
		})

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error message": "token is not valid"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			c.Set("claims", claims)
		}

		c.Next()
	}
}

func roleBasedMiddleware(desiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		yourRole := claims.(jwt.MapClaims)["role"].(string)

		if yourRole != desiredRole {
			c.JSON(http.StatusUnauthorized, gin.H{"error message": "You are not allowed to access this resource"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.POST("/login/admin", func(c *gin.Context) {
		username := "hariom679"
		userrole := "admin"
		tokenString, _ := tokenGenerator(username, userrole)
		c.JSON(200, gin.H{"Your token is ": tokenString})
	})

	router.POST("/login/user", func(c *gin.Context) {
		username := "hariom679"
		userrole := "user"
		tokenString, _ := tokenGenerator(username, userrole)
		c.JSON(200, gin.H{"Your token is ": tokenString})
	})

	protected := router.Group("/protected")
	protected.Use(authHandlerMiddleware())
	{
		protected.GET("/user", roleBasedMiddleware("user"), func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"Welcome": "Hello! user partner---"})
		})
		protected.GET("/admin", roleBasedMiddleware("admin"), func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"Welcome": "Hello! admin partner---"})
		})
	}

	router.Run(":9090")

}
