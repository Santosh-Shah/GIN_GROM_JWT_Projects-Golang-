//package models
//
//type User struct {
//	ID          uint   `json:"id" gorm:"primaryKey"`
//	Username    string `json:"username" gorm:"unique;not null"`
//	Password    string `json:"password" gorm:"not null"`
//	FirstName   string `json:"firstname" gorm:"not null"`
//	LastName    string `json:"lastname" gorm:"not null"`
//	Email       string `json:"email" gorm:"unique;not null"`
//	PhoneNumber string `json:"phonenumber" gorm:"not null"`
//	CreatedAt   int64  `json:"created_at" gorm:"autoCreateTime"`
//	UpdatedAt   int64  `json:"updated_at" gorm:"autoUpdateTime"`
//}

package models

type User struct {
	UserID      string `json:"user_id" gorm:"unique;not null"`
	Username    string `json:"username" gorm:"unique;not null"`
	Password    string `json:"password" gorm:"not null"`
	FirstName   string `json:"firstname" gorm:"not null"`
	LastName    string `json:"lastname" gorm:"not null"`
	Email       string `json:"email" gorm:"unique;not null"`
	PhoneNumber string `json:"phonenumber" gorm:"not null"`
	CreatedAt   int64  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   int64  `json:"updated_at" gorm:"autoUpdateTime"`
}
