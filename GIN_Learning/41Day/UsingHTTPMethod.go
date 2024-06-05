package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	// Define route handlers for different HTTP methods
	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	router.PATCH("/somePatch", patching)
	router.HEAD("/someHead", head)
	router.OPTIONS("/someOptions", options)

	// By default, it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run(":8080")
	// router.Run(":3000") for a hard-coded port
}

// Route handler for GET requests to "/someGet"
func getting(c *gin.Context) {
	c.String(http.StatusOK, "GET request received")
}

// Route handler for POST requests to "/somePost"
func posting(c *gin.Context) {
	c.String(http.StatusOK, "POST request received")
}

// Route handler for PUT requests to "/somePut"
func putting(c *gin.Context) {
	c.String(http.StatusOK, "PUT request received")
}

// Route handler for DELETE requests to "/someDelete"
func deleting(c *gin.Context) {
	c.String(http.StatusOK, "DELETE request received")
}

// Route handler for PATCH requests to "/somePatch"
func patching(c *gin.Context) {
	c.String(http.StatusOK, "PATCH request received")
}

// Route handler for HEAD requests to "/someHead"
func head(c *gin.Context) {
	c.String(http.StatusOK, "HEAD request received")
}

// Route handler for OPTIONS requests to "/someOptions"
func options(c *gin.Context) {
	c.String(http.StatusOK, "OPTIONS request received")
}
