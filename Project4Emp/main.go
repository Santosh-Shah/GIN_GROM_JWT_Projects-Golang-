package main

import (
	"SimplejwtProject/Project4Emp/controllers"
	"SimplejwtProject/Project4Emp/models"
	"SimplejwtProject/Project4Emp/router"
	"SimplejwtProject/Project4Emp/services"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@Mysql_679#@tcp(127.0.0.1:3306)/emp_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&models.Employee{})
	if err != nil {
		panic("failed to migrate database")
	}

	services.InitDB(db)

	router := gin.Default()

	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)

	protected := router.Group("/protected")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/employees", controllers.GetAllEmployees)
	protected.PUT("/employees/:id", controllers.UpdateEmployee)
	protected.DELETE("/employees/:id", controllers.DeleteEmployee)

	admin := router.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())

	admin.GET("/employees", controllers.GetAllEmployees)
	admin.PUT("/employees/:id", controllers.UpdateEmployee)
	admin.DELETE("/employees/:id", controllers.DeleteEmployee)

	router.Run(":8080")
}
