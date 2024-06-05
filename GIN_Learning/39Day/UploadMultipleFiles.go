package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Endpoint for uploading multiple files
	router.POST("/upload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()

		// Retrieve all files from the "upload[]" field
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename) // Log the filename of each uploaded file

			// Specify the destination path to save the uploaded file
			dst := "uploads/" + file.Filename

			// Save the uploaded file to the specified destination
			if err := c.SaveUploadedFile(file, dst); err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("error saving file: %s", err.Error()))
				return
			}
		}

		// Respond with a success message indicating the number of files uploaded
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	router.Run(":8080")
}
