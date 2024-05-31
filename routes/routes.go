package routes

import (
	"github.com/aleksanderpalamar/rbac-api/controllers"
	"github.com/aleksanderpalamar/rbac-api/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	router.GET("/users", middlewares.IsAuthorized(), controllers.GetUsers)
	router.POST("/users", controllers.CreateUser)
	router.POST("/login", controllers.Login)
}
