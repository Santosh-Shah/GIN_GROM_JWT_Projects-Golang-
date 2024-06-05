package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with default middleware
	router := gin.Default()

	// Define a route to serve data from a reader
	router.GET("/someDataFromReader", func(c *gin.Context) {
		// Send an HTTP GET request to fetch data from a remote URL
		response, err := http.Get("https://ugc.production.linktr.ee/55eee7a3-7d20-451f-986c-56295e5a61ca_SantoshShah.jpeg?io=true&size=avatar-v3_0")
		if err != nil || response.StatusCode != http.StatusOK {
			// If there is an error or the status code is not OK, return 503 Service Unavailable
			c.Status(http.StatusServiceUnavailable)
			return
		}

		// Retrieve the response body as a reader, content length, and content type
		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		// Define extra headers to be sent in the response
		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		// Serve data from the reader as the response body
		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})

	// Run the Gin server on port 8080
	router.Run(":8080")
}
