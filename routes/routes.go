package routes

import (
	"github.com/aleksanderpalamar/rbac-api/controllers"
	"github.com/aleksanderpalamar/rbac-api/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(router *gin.Engine) {
	router.GET("/users", middlewares.IsAuthorized(), controllers.GetUsers)
	router.POST("/users", controllers.CreateUser)
	router.POST("/login", controllers.Login)

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
