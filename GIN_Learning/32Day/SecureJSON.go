package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router with default middleware
	r := gin.Default()

	// Define a route for serving JSON data securely
	r.GET("/someJSON", func(c *gin.Context) {
		// Define some data to be sent as JSON
		names := []string{"lena", "austin", "foo"}

		// Send JSON response securely using SecureJSON method
		// This method prevents JSON hijacking by prepending a secure prefix
		c.SecureJSON(http.StatusOK, names)
	})

	// Listen and serve on port 8080
	r.Run(":8080")
}
