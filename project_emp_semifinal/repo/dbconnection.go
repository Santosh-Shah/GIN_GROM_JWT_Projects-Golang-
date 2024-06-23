package repo

import (
	"SimplejwtProject/project_emp_semifinal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	var err error
	dsn := "root:@Mysql_679#@tcp(127.0.0.1:3306)/emp_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.Employee{})
	if err != nil {
		panic(err)
	}
}
