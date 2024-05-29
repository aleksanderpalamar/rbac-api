package routes

import (
	"github.com/aleksanderpalamar/rbac-api/controllers"
	"github.com/aleksanderpalamar/rbac-api/middlewares"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users", middlewares.IsAuthorized("admin", controllers.CreateUser)).Methods("GET")

	return router
}
