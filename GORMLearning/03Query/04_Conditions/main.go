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

	//TODO: Retrieving all objects
	//TODO: sql query
	//SELECT * FROM albums;

	var albums []Album
	db.Find(&albums)
	for _, a := range albums {
		fmt.Println(a)
	}

}
