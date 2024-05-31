package controllers

import (
	"net/http"

	"github.com/aleksanderpalamar/rbac-api/auth"
	"github.com/aleksanderpalamar/rbac-api/models"
	"github.com/aleksanderpalamar/rbac-api/utils"
	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Failure 401 {object} gin.H{"error": "Unauthorized"}
// @Router /users [get]
func GetUsers(c *gin.Context) {
	var users []models.User

	if err := utils.DB.Preload("Role").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "User"
// @Success 200 {object} models.User
// @Failure 400 {object} gin.H{"error": "Bad Request"}
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, input)
}

// Login godoc
// @Summary Login user
// @Description Login user
// @Tags auth
// @Accept  json
// @Produce  json
// @Param login body models.User true "Login"
// @Success 200 {object} gin.H{"token": "token"}
// @Failure 401 {object} gin.H{"error": "Unauthorized"}
// @Router /login [post]
func Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := auth.GenerateJWT(input.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
