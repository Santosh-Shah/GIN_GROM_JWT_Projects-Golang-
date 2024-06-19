package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
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

		// Add token claims to the context for use in handlers
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("claims", claims)
		}

		c.Next()
	}
}

func roleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No claims found in context"})
			c.Abort()
			return
		}

		userRole := claims.(jwt.MapClaims)["role"].(string)
		if userRole != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't have access to this resource"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.POST("/login/admin", func(c *gin.Context) {
		username := "Hariom Shah"
		role := "admin"
		tokenString, err := generateJWT(username, role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	})

	r.POST("/login/user", func(c *gin.Context) {
		username := "Hariom Shah"
		role := "user"
		tokenString, err := generateJWT(username, role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	})

	protected := r.Group("/protected")
	protected.Use(authMiddleware())
	{
		protected.GET("/admin", roleMiddleware("admin"), func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Welcome, admin!"})
		})

		protected.GET("/user", roleMiddleware("user"), func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Welcome, user!"})
		})
	}

	r.Run(":8082")
}
