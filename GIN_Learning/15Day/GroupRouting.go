package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Handler functions for the endpoints
func loginEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "login successful"})
}

func submitEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "submit successful"})
}

func readEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "read successful"})
}

func main() {
	// Initialize the Gin router with default middleware (logger and recovery)
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	v2 := router.Group("v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	router.Run(":8080")
}
