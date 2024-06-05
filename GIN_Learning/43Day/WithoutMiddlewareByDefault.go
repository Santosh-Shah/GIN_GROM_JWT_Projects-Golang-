package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.New()

	// Default With the Logger and Recovery middleware already attached
	//r := gin.Default()
	fmt.Println(r)
}
