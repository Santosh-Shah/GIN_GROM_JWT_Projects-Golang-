package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//func main() {
//	r := gin.Default()
//	r.GET("/ping", func(c *gin.Context) {
//		c.IndentedJSON(http.StatusOK, gin.H{
//			"Golang Alert": "I am learning GIN framework",
//		})
//	})
//	r.Run() // listen and serve on 0.0.0.0:8080
//}

func main() {
	router := gin.Default()
	router.GET("/someJSON", func(context *gin.Context) {
		context.Header("Content-Type", "application/json")
		date := map[string]interface{}{
			"lang":        "GO语言",
			"tag":         "..!#",
			"age":         18,
			"isDeveloper": true,
		}

		context.AsciiJSON(http.StatusOK, date)
	})

	router.Run(":8080")
}
