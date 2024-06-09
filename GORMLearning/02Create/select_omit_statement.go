package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

//type User struct {
//	ID       uint `gorm:"primarykey"`
//	Name     string
//	Age      int
//	Birthday time.Time
//}

type User struct {
	ID       uint `gorm:"primary_key"`
	Name     string
	Age      int
	Birthday time.Time
}

func main() {
	dsn := "root:@Mysql_679#@tcp(127.0.0.1:3306)/santosh_snippat?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("Database connection failed: ", err)
	}

	// creating simple records
	user := User{Name: "Raghav Ghimire", Age: 75, Birthday: time.Now()}

	// creating multiple records
	//users := []*User{
	//	{Name: "Hariom Shah", Age: 22, Birthday: time.Now()},
	//	{Name: "Omprakash Shah", Age: 20, Birthday: time.Now()},
	//}

	// using select statement
	//db.Select("Name", "Birthday").Create(&user)

	//using omit statement
	db.Omit("Age").Create(&user)

	//result := db.Create(&user)

	//log.Println("User: ", user)
	//log.Println(result)
	////log.Println(user.ID)
	//log.Println(result.Error)
	//log.Println(result.RowsAffected)

}
