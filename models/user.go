package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password"`
	Role     string `json:"role" gorm:"foreignKey:RoleID"`
	RoleID   uint   `json:"role_id"`
}

type Role struct {
	gorm.Model
	Name string `json:"name" gorm:"unique;not null"`
}
