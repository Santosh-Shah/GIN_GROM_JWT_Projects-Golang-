package main

import "github.com/gin-gonic/gin"

// without colorizingNever colorize logs:c
//func main() {
//	gin.DisableConsoleColor()
//
//	// Creates a gin router with default middleware:
//	// logger and recovery (crash-free) middleware
//	router := gin.Default()
//
//	router.GET("/ping", func(c *gin.Context) {
//		c.String(200, "pong")
//	})
//
//	router.Run(":8080")
//}

//Always colorize logs:

func main() {
	// Force log's color
	gin.ForceConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware

	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.String(200, "pong")
	})

	router.Run(":8080")
}
