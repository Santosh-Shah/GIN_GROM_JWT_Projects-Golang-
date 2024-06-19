package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type Album struct {
	Id     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int64  `json:"price"`
}

var db *gorm.DB

func initDB() {
	var err error
	dsn := "root:@Mysql_679#@tcp(127.0.0.1:3306)/album_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	err = db.AutoMigrate(&Album{})
	if err != nil {
		panic("Failed to migrate database")
	}
}

func getAllAlbums(c *gin.Context) {
	var albums []Album
	db.Find(&albums)
	c.IndentedJSON(http.StatusOK, gin.H{"albums": albums})
}

func getAlbumById(c *gin.Context) {
	var album Album
	id := c.Param("id")
	result := db.First(&album, id)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"album": album})
}

func createAlbum(c *gin.Context) {
	var album Album
	err := c.BindJSON(&album)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&album)
	c.IndentedJSON(http.StatusCreated, album)
}

//func createAlbum(context *gin.Context) {
//	var newAlbum Album
//	if err := context.BindJSON(&newAlbum); err != nil {
//		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
//		return
//	}
//	db.Create(&newAlbum)
//	context.IndentedJSON(http.StatusCreated, newAlbum)
//}

func deleteAlbumById(c *gin.Context) {
	var album Album
	id := c.Param("id")
	result := db.First(&album, id)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
		return
	}

	db.Delete(&album)
	c.IndentedJSON(http.StatusOK, gin.H{"album": album})
}

func main() {
	initDB()
	router := gin.Default()
	router.GET("/albums", getAllAlbums)
	router.GET("/album/:id", getAlbumById)
	router.POST("/album", createAlbum)
	router.DELETE("/album/:id", deleteAlbumById)
	router.Run(":8080")
}
