package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//func main() {
//	router := gin.Default()
//
//	// Use http.ListenAndServe to start the server
//	http.ListenAndServe(":8080", router)
//}

func main() {
	router := gin.Default()

	// Create a http.Server instance with custom settings
	s := &http.Server{
		Addr:           ":8080",          // Address to listen on
		Handler:        router,           // Handler to invoke (Gin router)
		ReadTimeout:    10 * time.Second, // Maximum duration for reading the entire request, including the body
		WriteTimeout:   10 * time.Second, // Maximum duration before timing out writes of the response
		MaxHeaderBytes: 1 << 20,          // Maximum size of request headers (1 MB)
	}

	// Start the server
	s.ListenAndServe()
}
