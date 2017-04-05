package models

import (
	"github.com/jinzhu/gorm"
)

// User Model
type User struct {
	gorm.Model
	Name      string `json:"name"`
	Title     string `json:"title"`
	Email     string `gorm:"unique_index" json:"email" valid:"required,email"`
	Password  string `json:"password" valid:"required"`
	Avatar    string `json:"avatar"`
	Contact   string `json:"contact"`
	Address   string `json:"address"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Token     string `json:"token"`
}
