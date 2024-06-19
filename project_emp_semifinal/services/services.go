package services

import (
	"SimplejwtProject/project_emp_semifinal/models"
	"errors"
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

func CreateEmployee(emp *models.Employee) {
	db.Create(emp)
}

func GetAllEmployee() []models.Employee {
	var emps []models.Employee
	db.Find(&emps)
	return emps
}

func GetEmployeesById(id string) *models.Employee {
	var emp *models.Employee
	db.First(&emp, id)
	return emp
}

func DeleteEmployeeById(id string) models.Employee {
	var emp models.Employee
	db.First(&emp, id)
	db.Delete(&emp)
	return emp
}

func UpdateEmployeeById(emp *models.Employee) error {
	return db.Save(emp).Error
}

func CheckUsernameOrEmailExists(username string, email string) error {
	var count int64
	db.Model(&models.Employee{}).Where("username = ? OR email = ?", username, email).Count(&count)
	if count > 0 {
		return errors.New("employee with username or email already exists")
	}

	return nil
}

func GetEmployeeByUsername(username string) *models.Employee {
	var emp models.Employee
	db.Where("username = ?", username).First(&emp)
	return &emp
}
