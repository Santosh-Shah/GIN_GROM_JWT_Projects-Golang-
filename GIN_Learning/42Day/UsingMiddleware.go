package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// SimpleMiddleware is a middleware that logs the incoming request method and path.
func SimpleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("Incoming %s request to %s\n", c.Request.Method, c.Request.URL.Path)
		c.Next() // Continue to the next middleware/handler
	}
}

func main() {
	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Per route middleware
	r.GET("/benchmark", SimpleMiddleware(), benchEndpoint)

	// Authorization group
	authorized := r.Group("/")
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", loginEndpoint)
		authorized.POST("/submit", submitEndpoint)
		authorized.POST("/read", readEndpoint)

		// Nested group
		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

// AuthRequired is a dummy authentication middleware for learning purposes.
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Dummy authentication logic for learning purposes
		fmt.Println("Authentication logic goes here (for learning purposes)")
		c.Next()
	}
}

// Placeholder handler functions for endpoints

func benchEndpoint(c *gin.Context) {
	c.String(200, "Benchmark endpoint")
}

func loginEndpoint(c *gin.Context) {
	c.String(200, "Login endpoint")
}

func submitEndpoint(c *gin.Context) {
	c.String(200, "Submit endpoint")
}

func readEndpoint(c *gin.Context) {
	c.String(200, "Read endpoint")
}

func analyticsEndpoint(c *gin.Context) {
	c.String(200, "Analytics endpoint")
}
