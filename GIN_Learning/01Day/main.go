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
	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "hello",
		}

		// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
