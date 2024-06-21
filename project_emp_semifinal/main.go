package main

import (
	"SimplejwtProject/project_emp_semifinal/controllers"
	middleware "SimplejwtProject/project_emp_semifinal/router"
	"SimplejwtProject/project_emp_semifinal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	services.InitDB()

	router := gin.Default()
	router.POST("/signup", controllers.Signup)
	router.POST("login", controllers.Login)

	router.GET("/employees", controllers.GetAllEmployees)
	router.GET("/employee/id/:id", controllers.GetEmployeesById)
	router.GET("/employee/username/:username", controllers.GetEmployeeByUsername)
	router.DELETE("/employee/:id", controllers.DeleteEmployeeById)
	router.PUT("/employee/:id", controllers.UpdateEmployeeById)

	protected := router.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/employees", controllers.GetAllEmployees)
		protected.GET("/employee/id/:id", controllers.GetEmployeesById)
		protected.GET("/employee/username/:username", controllers.GetEmployeeByUsername)
		protected.DELETE("/employee/:id", controllers.DeleteEmployeeById)
		protected.PUT("/employee/:id", controllers.UpdateEmployeeById)
	}

	admin := router.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	{
		admin.GET("/employees", controllers.GetAllEmployees)
		admin.GET("/employee/id/:id", controllers.GetEmployeesById)
		admin.GET("/employee/username/:username", controllers.GetEmployeeByUsername)
		admin.DELETE("/employee/:id", controllers.DeleteEmployeeById)
		admin.PUT("/employee/:id", controllers.UpdateEmployeeById)

		admin.GET("/profile", controllers.GetProfileByToken)
	}

	user := router.Group("/user")
	user.Use(middleware.AuthMiddleware())
	{
		user.GET("/profile", controllers.GetUserProfile)
	}

	err := router.Run(":8080")
	if err != nil {
		return
	}

}
