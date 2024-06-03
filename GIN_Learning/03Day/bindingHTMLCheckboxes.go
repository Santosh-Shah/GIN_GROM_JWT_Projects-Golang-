package main

import (
	"github.com/gin-gonic/gin"
)

type myForm struct {
	Colors []string `form:"colors[]"`
}

func formHandler(c *gin.Context) {
	var fakeForm myForm
	c.ShouldBind(&fakeForm)
	c.JSON(200, gin.H{"color": fakeForm.Colors})
}

func main() {
	r := gin.Default()

	// Serve static files from the templates directory
	r.Static("/static", "GIN_Learning/03Day/templates/form.html")

	// Define route handlers
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "form.html", nil)
	})

	r.POST("/submit", formHandler)

	// Run the server
	r.Run(":8080")
}
