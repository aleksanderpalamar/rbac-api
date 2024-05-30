package main

import (
	"log"
	"net/http"

	"github.com/aleksanderpalamar/rbac-api/routes"
	"github.com/aleksanderpalamar/rbac-api/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	utils.ConnectionDB()

	router := mux.NewRouter()
	routes.SetupRouter(router)

	log.Println("Server started on port 3001")
	log.Fatal(http.ListenAndServe(":3001", router))
}
