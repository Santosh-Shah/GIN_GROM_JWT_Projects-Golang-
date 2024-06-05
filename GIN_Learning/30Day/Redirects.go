package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// GET Redirect to External Location
	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://linktr.ee/santosh_shah")
	})

	// POST Redirect to Internal Location
	r.POST("/test-internal", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/foo")
	})

	// Router Redirect
	r.GET("/test-internal", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})

	// Start the server
	r.Run(":8080")
}
