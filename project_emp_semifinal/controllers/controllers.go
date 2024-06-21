package controllers

import (
	"SimplejwtProject/project_emp_semifinal/models"
	"SimplejwtProject/project_emp_semifinal/router"
	"SimplejwtProject/project_emp_semifinal/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Signup(c *gin.Context) {
	var emp models.Employee

	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CheckUsernameOrEmailExists(emp.Username, emp.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(emp.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	emp.Password = string(hashPassword)

	services.CreateEmployee(&emp)
	c.IndentedJSON(http.StatusCreated, emp)
}

func GetAllEmployees(c *gin.Context) {
	emp := services.GetAllEmployee()
	c.IndentedJSON(http.StatusOK, emp)

}

func GetEmployeesById(c *gin.Context) {
	id := c.Param("id")
	emp := services.GetEmployeesById(id)
	if emp == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
	}

	c.IndentedJSON(http.StatusOK, emp)
}

func UpdateEmployeeById(c *gin.Context) {
	id := c.Param("id")
	var input models.Employee

	// Bind the incoming JSON to the input variable
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the existing employee by ID
	emp := services.GetEmployeesById(id)
	//if emp == nil {
	//	c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
	//	return
	//}

	// Update the fields with the input data
	emp.FirstName = input.FirstName
	emp.LastName = input.LastName
	emp.Email = input.Email
	emp.PhoneNumber = input.PhoneNumber

	// Save the updated employee record
	err := services.UpdateEmployeeById(emp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update employee"})
		return
	}

	c.IndentedJSON(http.StatusOK, emp)
}

func DeleteEmployeeById(c *gin.Context) {
	id := c.Param("id")
	emp := services.DeleteEmployeeById(id)
	c.IndentedJSON(http.StatusOK, emp)
}

func GetEmployeeByUsername(c *gin.Context) {
	username := c.Param("username")
	emp := services.GetEmployeeByUsername(username)
	c.IndentedJSON(http.StatusOK, emp)
}

func Login(c *gin.Context) {
	type Login struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var logEmp Login

	if err := c.ShouldBindJSON(&logEmp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	emp := services.GetEmployeeByUsername(logEmp.Username)

	err := bcrypt.CompareHashAndPassword([]byte(emp.Password), []byte(logEmp.Password))

	if logEmp.Username != emp.Username && err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	signedString, err := middleware.TokenGenerator(logEmp.Username, logEmp.Password, emp.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"Your Token": signedString})
}

func GetUserProfile(c *gin.Context) {
	claims, _ := c.Get("claims")
	role := claims.(jwt.MapClaims)["role"].(string)
	if role == "admin" {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"Unauthorized": "You have no authority to access admin profile"})
		c.Abort()
		return
	}

	username := claims.(jwt.MapClaims)["username"].(string)
	emp := services.GetEmployeeByUsername(username)
	c.IndentedJSON(http.StatusOK, emp)
}

func UpdateUserProfile(c *gin.Context) {

}

func UpdateAdminProfile(c *gin.Context) {

}

func GetProfileByToken(c *gin.Context) {
	claims, _ := c.Get("claims")
	username := claims.(jwt.MapClaims)["username"].(string)
	emp := services.GetEmployeeByUsername(username)
	c.IndentedJSON(http.StatusOK, emp)
}
