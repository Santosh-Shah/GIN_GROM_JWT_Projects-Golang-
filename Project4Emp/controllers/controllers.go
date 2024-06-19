package controllers

import (
	middleware "SimplejwtProject/Project4Emp/router"
	"net/http"
	"strconv"

	"SimplejwtProject/Project4Emp/models"
	"SimplejwtProject/Project4Emp/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var input models.Employee
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CheckUsernameOrEmailExists(input.Username, input.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	input.Password = string(hashedPassword)
	if err := services.CreateEmployee(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create employee"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee created successfully"})
}

func Login(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee, err := services.GetEmployeeByUsername(input.Username)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(input.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	tokenString, err := middleware.GenerateJWT(employee.UserID, string(employee.Role))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func GetAllEmployees(c *gin.Context) {
	employees, err := services.GetAllEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get employees"})
		return
	}
	c.JSON(http.StatusOK, employees)
}

func UpdateEmployee(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input models.Employee
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee, err := services.GetEmployeeByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	employee.FirstName = input.FirstName
	employee.LastName = input.LastName
	employee.Email = input.Email
	employee.PhoneNumber = input.PhoneNumber

	if err := services.UpdateEmployee(&employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update employee"})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func DeleteEmployee(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := services.DeleteEmployee(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete employee"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}
