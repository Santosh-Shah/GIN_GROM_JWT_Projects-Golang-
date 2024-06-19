package services

import (
	"SimplejwtProject/Project4Emp/models"
	"errors"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(database *gorm.DB) {
	db = database
}

func CheckUsernameOrEmailExists(username, email string) error {
	var count int64
	db.Model(&models.Employee{}).Where("username = ? OR email = ?", username, email).Count(&count)
	if count > 0 {
		return errors.New("username or email already exists")
	}
	return nil
}

func CreateEmployee(employee *models.Employee) error {
	return db.Create(employee).Error
}

func GetEmployeeByUsername(username string) (models.Employee, error) {
	var employee models.Employee
	err := db.Where("username = ?", username).First(&employee).Error
	return employee, err
}

func GetEmployeeByID(userID uint) (models.Employee, error) {
	var employee models.Employee
	err := db.First(&employee, userID).Error
	return employee, err
}

func UpdateEmployee(employee *models.Employee) error {
	return db.Save(employee).Error
}

func DeleteEmployee(userID uint) error {
	return db.Delete(&models.Employee{}, userID).Error
}

func GetAllEmployees() ([]models.Employee, error) {
	var employees []models.Employee
	err := db.Find(&employees).Error
	return employees, err
}
