package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

// router01 defines a Gin router for server 01.
func router01() http.Handler {
	// Create a new Gin engine
	e := gin.New()
	// Use recovery middleware to recover from panics
	e.Use(gin.Recovery())
	// Define a route that responds with JSON data "Welcome server 01" when accessed
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":    http.StatusOK,
				"message": "Welcome server 01",
			},
		)
	})
	return e
}

// router02 defines a Gin router for server 02.
func router02() http.Handler {
	// Create a new Gin engine
	e := gin.New()
	// Use recovery middleware to recover from panics
	e.Use(gin.Recovery())
	// Define a route that responds with JSON data "Welcome server 02" when accessed
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":    http.StatusOK,
				"message": "Welcome server 02",
			},
		)
	})
	return e
}

func main() {
	// Create server01 with its address, handler, read and write timeouts
	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Create server02 with its address, handler, read and write timeouts
	server02 := &http.Server{
		Addr:         ":8081",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start server01 in a separate goroutine
	g.Go(func() error {
		return server01.ListenAndServe()
	})

	// Start server02 in a separate goroutine
	g.Go(func() error {
		return server02.ListenAndServe()
	})

	// Wait for both servers to finish and log any errors
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
