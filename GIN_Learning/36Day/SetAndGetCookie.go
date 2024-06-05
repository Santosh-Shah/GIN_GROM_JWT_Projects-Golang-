package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with default middleware
	router := gin.Default()

	// Define a route to handle GET requests to /cookie
	router.GET("/cookie", func(c *gin.Context) {
		// Attempt to retrieve the value of the cookie named "gin_cookie"
		cookie, err := c.Cookie("gin_cookie")

		// Check if an error occurred while retrieving the cookie
		if err != nil {
			// If the cookie is not set, initialize it with a default value and set it in the response
			cookie = "NotSet"
			// SetCookie method sets the cookie in the response
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}

		// Print the value of the cookie
		fmt.Printf("Cookie value: %s \n", cookie)
	})

	// Run the Gin server
	router.Run()
}
