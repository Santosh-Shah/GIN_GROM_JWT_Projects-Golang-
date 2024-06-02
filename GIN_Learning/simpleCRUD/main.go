package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Album struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int64  `json:"price"`
}

var albums = []Album{
	{"1", "Mero Desh", "Laxmi Prasad Devkota", 400},
	{"2", "Mero Mount", "Climbers", 500},
	{"3", "The Terai", "Desi Author", 6000},
}

func getAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
	//context.JSON(http.StatusOK, albums)
}

func getAlbumById(context *gin.Context) {
	id := context.Param("id")

	for _, actualAlbum := range albums {
		if actualAlbum.Id == id {
			context.IndentedJSON(http.StatusOK, actualAlbum)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "sorry album not found! please try again"})
}

func postAlbum(context *gin.Context) {
	var newAlbum Album

	err := context.BindJSON(&newAlbum)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request body! please try again"})
		return
	}

	albums = append(albums, newAlbum)
	context.IndentedJSON(http.StatusCreated, newAlbum)
}

func deleteAlbumById(context *gin.Context) {
	id := context.Param("id")

	for index, dAlbum := range albums {
		if dAlbum.Id == id {
			albums = append(albums[:index], albums[index+1:]...)
			context.IndentedJSON(http.StatusOK, dAlbum)
			return
		}
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found! please try again"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/getAlbum/:id", getAlbumById)
	router.POST("/album", postAlbum)
	router.DELETE("/album/:id", deleteAlbumById)
	router.Run(":8080")
}
