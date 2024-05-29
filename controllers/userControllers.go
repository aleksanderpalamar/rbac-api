package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/aleksanderpalamar/rbac-api/models"
	"github.com/aleksanderpalamar/rbac-api/utils"
	"gorm.io/gorm"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := utils.DB.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	if err := utils.DB.Preload("Role").Find(&users).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "No users found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(users)
}
