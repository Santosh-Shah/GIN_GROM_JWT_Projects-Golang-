package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	// Create a default Gin router
	r := gin.Default()

	// Define a simple ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Set up a custom autocert manager
	m := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example.com"),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}

	// Create a HTTPS server with custom TLS configuration
	server := &http.Server{
		Addr:    ":443",
		Handler: r,
		TLSConfig: &tls.Config{
			GetCertificate: m.GetCertificate,
		},
	}

	// Run the HTTPS server with automatic certificate management
	log.Fatal(server.ListenAndServeTLS("", ""))
}
