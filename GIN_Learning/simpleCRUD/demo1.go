package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Albums struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int64  `json:"price"`
}

var album = []Albums{
	{"1", "Mero Desh", "Laxmi Prasad Devkota", 400},
	{"2", "Mero Mount", "Climbers", 500},
	{"3", "The Terai", "Desi Author", 6000},
}

func postAlbums(c *gin.Context) {
	var newAlbum Albums
	err := c.BindJSON(&newAlbum)
	if err != nil {
		c.JSON(400, gin.H{"message": "album creation failed"})
	}

	album = append(album, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)

}

func getAlbumsById(c *gin.Context) {
	id := c.Param("id")
	for _, album := range album {
		if album.Id == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func getAllAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, album)
}

func deleteAlbumsById(c *gin.Context) {
	id := c.Param("id")
	for index, a := range album {
		if a.Id == id {
			album = append(album[:index], album[index+1:]...)
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found! please try again"})
}

func main() {
	router := gin.Default()
	router.POST("/album", postAlbums)
	router.GET("/getAlbum/:id", getAlbumsById)
	router.GET("/getAllAlbums", getAllAlbums)
	router.DELETE("/deleteAlbum/:id", deleteAlbumsById)
	router.Run(":9090")
}
