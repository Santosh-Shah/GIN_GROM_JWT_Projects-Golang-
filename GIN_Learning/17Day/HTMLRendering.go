package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Initialize the Gin router with default middleware
	router := gin.Default()

	// Load HTML templates from the "templates" directory using LoadHTMLGlob
	router.LoadHTMLGlob("templates/*")
	// Alternatively, you can specify individual HTML files using LoadHTMLFiles
	// router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

	// Define a route to render the index page
	router.GET("/index", func(c *gin.Context) {
		// Render the "index.tmpl" template with a title
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	// Start the server on port 8080
	router.Run(":8080")
}
