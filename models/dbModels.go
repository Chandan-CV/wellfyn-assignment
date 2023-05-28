package models

import (
	"gorm.io/gorm"
)

// this is the model for the table userauths in the database
type Userauth struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique not null"`
	Password string `json:"password" gorm:"not null"`
}

// this is the model for the table user_data in the database
type UserData struct {
	gorm.Model
	Email     string `json:"email" gorm:"unique not null"`
	Name      string `json:"name" gorm:"not null"`
	Phone     string `json:"phone" gorm:"not null"`
	Images    string `json:"images"`
	Documents string `json:"documents"`
}
