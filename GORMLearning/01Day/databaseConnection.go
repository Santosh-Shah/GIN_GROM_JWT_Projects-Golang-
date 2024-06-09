package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

//func main() {
//	// Data Source Name (DSN) contains the database connection details
//	dsn := "root:@Mysql_679#@tcp(127.0.0.1:3306)/santosh_snippat?charset=utf8mb4&parseTime=True&loc=Local"

//	// Initialize a GORM DB connection using the MySQL dialector
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//
//	// Check if there was an error connecting to the database
//	if err != nil {
//		log.Fatal("Failed to connect to database: ", err)
//	}
//
//	// Use the 'db' object to interact with the database
//	log.Println("Congratulations database connected successfully!", db)
//}

func main() {
	dsn := "root:@Mysql_679#@tcp(127.0.0.1:3306)/santosh_snippat?charset=utf8mb4&parseTime=True&loc=Local"

	dia := mysql.Open(dsn)

	db, err := gorm.Open(dia, &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Congratulations database is connected successfully!", db)
}
