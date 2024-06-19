package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"SimplejwtProject/Project4Emp/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("your_secret_key")

func GenerateJWT(userID uint, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userID
	claims["Role"] = role
	claims["exp_time"] = time.Now().Add(time.Hour * 24).Unix()
	return token.SignedString(JwtKey)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		if !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Print out the token for debugging
		fmt.Println("Token:", tokenString)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return JwtKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			c.Set("claims", claims)
		}

		c.Next()

		//claims, ok := token.Claims.(jwt.MapClaims)
		//if ok && token.Valid {
		//	c.Set("userID", uint(claims["userID"].(float64)))
		//	c.Set("role", claims["Role"].(string))
		//} else {
		//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		//	c.Abort()
		//	return
		//}
		//
		//c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != string(models.Admin) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}
		c.Next()
	}
}
