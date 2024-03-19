package model

import "gorm.io/gorm"

// User Model
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
