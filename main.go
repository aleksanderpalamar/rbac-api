package main

import (
	"log"
	"net/http"

	"github.com/aleksanderpalamar/rbac-api/routes"
	"github.com/aleksanderpalamar/rbac-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	utils.ConnectionDB()

	router := gin.Default()
	routes.SetupRouter(router)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)
	http.ListenAndServe(":3000", handler)
}
