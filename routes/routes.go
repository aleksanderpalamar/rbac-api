package routes

import (
	"github.com/aleksanderpalamar/rbac-api/controllers"
	"github.com/aleksanderpalamar/rbac-api/middlewares"

	"github.com/gorilla/mux"
)

func SetupRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users", middlewares.IsAuthorized(controllers.GetUsers)).Methods("GET")

	return router
}
