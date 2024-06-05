package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Endpoint for uploading a single file
	router.POST("/upload", func(c *gin.Context) {
		// Multipart form
		//form, _ := c.MultipartForm()

		// Retrieve the file from the form
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("error parsing form: %s", err.Error()))
			return
		}

		// Generate a safe filename for saving
		filename := filepath.Base(file.Filename)

		// Specify the destination path to save the uploaded file
		dst := "uploads/" + filename

		// Save the uploaded file to the specified destination
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error saving file: %s", err.Error()))
			return
		}

		// Respond with a success message indicating the filename and upload status
		c.String(http.StatusOK, fmt.Sprintf("File '%s' uploaded successfully", filename))
	})

	router.Run(":8080")
}
