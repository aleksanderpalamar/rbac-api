package main

import (
	"log"
	"os"

	_ "github.com/aleksanderpalamar/rbac-api/docs"
	"github.com/aleksanderpalamar/rbac-api/routes"
	"github.com/aleksanderpalamar/rbac-api/utils"
	"github.com/gin-gonic/gin"
)

// @title RBAC API
// @version 1.0
// @description This is a simple Role-Based Access Control (RBAC) API built with Go.
// @contact.name Aleksander Palamar
// @contact.url https://aleksanderpalamar.dev
// @license.name MIT
// @license.url https://github.com/aleksanderpalamar/rbac-api/blob/main/LICENSE
// @host localhost:3000
// @BasePath /
func main() {
	utils.LoadEnvVariables()
	utils.ConnectionDB()
	utils.Cors()

	router := gin.Default()
	routes.SetupRouter(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server", err)
	}
}
