package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Album struct {
	Id     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title" gorm:"default: unknown title"`
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

	//TODO: storing multiple records
	// SQL query:
	//INSERT INTO albums (title, author, price) VALUES
	//('Raga', 'Ram', 500),
	//('Sangeet', 'Shyam', 600);

	//albums := []*Album{
	//	{Title: "Title4", Author: "Author4", Price: 100},
	//	{Title: "Title5", Author: "Author5", Price: 200},
	//}

	//TODO: creating records with selected fields
	// SQL query: INSERT INTO albums(title, author) VALUES ("Raga", "Ram");

	//album := Album{Title: "Title6", Author: "Author6"}
	//result := db.Create(&albums)

	//db.Select("Title", "Author").Create(&album)

	//TODO: creating an album record from Map
	db.Model(&Album{}).Create(map[string]interface{}{
		"Title":  "Title10",
		"Author": "Author10",
		"Price":  80,
	})

	//fmt.Println(albums.Id)
	//fmt.Println(result.Error)
	//fmt.Println(result.RowsAffected)
}
