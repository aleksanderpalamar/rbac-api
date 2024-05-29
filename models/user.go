package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
	Email    string `gorm:"unique" json:"email"`
	Role     string `json:"role" gorm:"foreignKey:RoleID"`
	RoleID   uint   `json:"role_id"`
}

type Role struct {
	gorm.Model
	Name string `json:"name"`
}
