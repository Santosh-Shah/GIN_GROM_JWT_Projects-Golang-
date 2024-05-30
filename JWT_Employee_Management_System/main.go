package main

import (
	"SimplejwtProject/JWT_Employee_Management_System/controllers"
	"SimplejwtProject/JWT_Employee_Management_System/database"
	"SimplejwtProject/JWT_Employee_Management_System/middleware"
	"SimplejwtProject/JWT_Employee_Management_System/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Initialize the database connection
	dsn := "root:@Mysql_679#@tcp(127.0.0.1:3306)/jwtproject1?charset=utf8mb4&parseTime=True&loc=Local"
	if err := database.Connect(dsn); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	if err := database.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Set up Gin router
	router := gin.Default()
	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)

	// Apply the AuthMiddleware to the routes you want to protect
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/users", controllers.GetUsers)
		protected.GET("/users/:id", controllers.GetUser)
		protected.PUT("/users/:id", controllers.UpdateUser)
		protected.DELETE("/users/:id", controllers.DeleteUser)
	}

	// Run the application
	port := "8080"
	if err := router.Run(":" + port); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
