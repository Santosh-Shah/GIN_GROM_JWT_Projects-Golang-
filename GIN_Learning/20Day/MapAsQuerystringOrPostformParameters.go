package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		// Parse query parameters as a map
		ids := c.QueryMap("ids")
		// Parse form parameters as a map
		names := c.PostFormMap("names")

		// Print the parsed maps
		fmt.Printf("ids: %v; names: %v\n", ids, names)

		// Respond with the parsed maps
		c.JSON(200, gin.H{
			"ids":   ids,
			"names": names,
		})
	})

	router.Run(":8080")
}
