package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/long_async", func(c *gin.Context) {
		// create copy to be used inside the goroutine
		cCp := c.Copy()
		go func() {
			// simulate a long task with time.Sleep(). 5 seconds
			time.Sleep(5 * time.Second)

			// note that you are using the copied context "cCp", IMPORTANT
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
		log.Println("from Goroutine-----------")
	})

	r.GET("/long_sync", func(c *gin.Context) {

		func() {
			// simulate a long task with time.Sleep(). 5 seconds
			time.Sleep(5 * time.Second)

			// since we are NOT using a goroutine, we do not have to copy the context
			log.Println("Done! in path " + c.Request.URL.Path)
		}()
		//// simulate a long task with time.Sleep(). 5 seconds
		//time.Sleep(5 * time.Second)
		log.Println("from Goroutine-----------")

	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
