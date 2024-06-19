package models

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	Admin Role = "admin"
	User  Role = "user"
)

type Employee struct {
	UserID      uint   `gorm:"primaryKey"`
	Username    string `gorm:"unique;not null"`
	Password    string `gorm:"not null"`
	FirstName   string `gorm:"not null"`
	LastName    string `gorm:"not null"`
	Email       string `gorm:"unique;not null"`
	PhoneNumber string `gorm:"not null"`
	Role        Role   `gorm:"not null;default:'user'"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
