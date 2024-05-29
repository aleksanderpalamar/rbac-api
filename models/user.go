package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	RoleID   uint   `json:"role_id"`
	Role     Role   `json:"role" gorm:"foreignKey:RoleID"`
}

type Role struct {
	gorm.Model
	Name  string `json:"name" gorm:"unique"`
	Users []User `gorm:"foreignKey:RoleID"`
}
