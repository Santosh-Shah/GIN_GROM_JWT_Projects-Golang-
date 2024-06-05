package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	// Disable Console Color, as it's not needed when writing logs to a file.
	gin.DisableConsoleColor()

	// Create a log file named "gin.log"
	f, err := os.Create("gin.log")
	if err != nil {
		panic(err)
	}
	// Set Gin's default writer to write to the file
	//gin.DefaultWriter = io.MultiWriter(f)

	// If you need to write logs to both file and console, use the following line instead:
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// Initialize the Gin router with default middleware (logger and recovery)
	router := gin.Default()

	// Define a simple GET endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Start the server on port 8080
	router.Run(":8080")
}
