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

	//TODO: get the first record order by primary key
	var albumOutput Album
	//db.First(&albumOutput)
	//fmt.Println("First user: ", albumOutput)

	//TODO: get any record, no specified order
	//db.Take(&albumOutput)
	//fmt.Println("Taken albumOutput:", albumOutput)

	//TODO: get last record ordered by primary key
	db.Last(&albumOutput)
	fmt.Println("Last Record: ", albumOutput)

}
