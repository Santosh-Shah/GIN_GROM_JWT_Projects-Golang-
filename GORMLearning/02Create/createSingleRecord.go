package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	ID       uint `gorm:"primarykey"`
	Name     string
	Age      int
	Birthday time.Time
}

func main() {
	// Data Source Name (DSN) contains the database connection details
	dsn := "root:@Mysql_679#@tcp(127.0.0.1:3306)/santosh_snippat?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	//initializing struct
	user := User{Name: "Santosh Shah", Age: 20, Birthday: time.Now()}

	// send data to mysql
	result := db.Create(&user)

	log.Println("User: ", user)
	log.Println(result)
	log.Println(user.ID)
	log.Println(result.Error)
	log.Println(result.RowsAffected)
}
