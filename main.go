package main

import (
	"log"
	"net/http"

	"github.com/aleksanderpalamar/rbac-api/routes"
	"github.com/aleksanderpalamar/rbac-api/utils"
	"github.com/gorilla/mux"
)

func main() {
	utils.ConnectionDB()

	router := mux.NewRouter()
	routes.SetupRouter(router)

	log.Println("Server started on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
