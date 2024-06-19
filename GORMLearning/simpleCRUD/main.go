package main

//
//import (
//	"net/http"
//
//	"github.com/gin-gonic/gin"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//)
//
//type Album struct {
//	Id     uint   `json:"id" gorm:"primaryKey"`
//	Title  string `json:"title"`
//	Author string `json:"author"`
//	Price  int64  `json:"price"`
//}
//
//var db *gorm.DB
//
//func initDB() {
//	var err error
//	dsn := "root:@Mysql_679#@tcp(127.0.0.1:3306)/album_db?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic("Failed to connect to database")
//	}
//
//	// Auto migrate will create the table. You can also handle table creation manually.
//	err = db.AutoMigrate(&Album{})
//	if err != nil {
//		panic("Failed to migrate database")
//	}
//}
//
//func getAlbums(context *gin.Context) {
//	var albums []Album
//	db.Find(&albums)
//	context.IndentedJSON(http.StatusOK, albums)
//}
//
//func getAlbumById(context *gin.Context) {
//	var album Album
//	id := context.Param("id")
//	if err := db.First(&album, id).Error; err != nil {
//		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
//		return
//	}
//	context.IndentedJSON(http.StatusOK, album)
//}
//
//func postAlbum(context *gin.Context) {
//	var newAlbum Album
//	if err := context.BindJSON(&newAlbum); err != nil {
//		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
//		return
//	}
//	db.Create(&newAlbum)
//	context.IndentedJSON(http.StatusCreated, newAlbum)
//}
//
//func deleteAlbumById(context *gin.Context) {
//	var album Album
//	id := context.Param("id")
//	if err := db.First(&album, id).Error; err != nil {
//		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
//		return
//	}
//	db.Delete(&album)
//	context.IndentedJSON(http.StatusOK, album)
//}
//
//func main() {
//	// Initialize database connection
//	initDB()
//
//	// Set up Gin router
//	router := gin.Default()
//
//	// Define routes
//	router.GET("/albums", getAlbums)
//	router.GET("/album/:id", getAlbumById)
//	router.POST("/album", postAlbum)
//	router.DELETE("/album/:id", deleteAlbumById)
//
//	// Start server
//	router.Run(":8080")
//}
