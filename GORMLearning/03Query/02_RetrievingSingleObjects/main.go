package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Album struct {
	Id     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int64  `json:"price"`
}

func main() {
	dsn := "root:@Mysql_679#@tcp(127.0.0.1:3306)/album_db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect with database")
		return
	}

	err = db.AutoMigrate(&Album{})
	if err != nil {
		panic("Failed to auto migrate")
		return
	}

	//TODO: storing simple records
	//albums := Album{Title: "Raga", Author: "Ram", Price: 4500}
	//db.Create(&albums)

	//TODO: Retrieve a single record by primary key (numeric)
	//TODO: sql query
	//SELECT * FROM users WHERE id = 10;

	//var album Album
	//db.First(&album, 9)
	//fmt.Println("Album", album)

	//db.First(&album, "5")
	//fmt.Println("Album", album)

	//TODO: Retrieve multiple records by primary key (numeric)
	//TODO: sql query
	//SELECT * FROM users WHERE id IN (1, 2, 3);

	//var albumOutput []Album
	//db.Find(&albumOutput, []int{3, 5, 6})
	//fmt.Println("Album", albumOutput)

	//var album = Album{Id: 9}
	//db.First(&album)
	//fmt.Println("Album", album)

	var album Album
	db.Model(Album{Id: 10}).First(&album)
	fmt.Println("Album", album)

	//TODO: Retrieve a single record by primary key (string/UUID)
	//TODO: sql query
	//SELECT * FROM users WHERE id = '1b74413f-f3b8-409f-ac47-e8c062e3472a';

	//TODO: Retrieve a single record when the object has a primary key value
	//TODO: sql query
	//SELECT * FROM users WHERE id = 10;

}
