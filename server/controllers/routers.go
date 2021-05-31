package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ogustavobelo/simple-crud-go/core"
	"github.com/ogustavobelo/simple-crud-go/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	BasePath = "simple-crud"
	Version  = "v1"
)

func InitRoutes(router *gin.Engine) {
	main := router.Group(BasePath)
	{
		// Swagger Setup
		docs.SwaggerInfo.Title = "Simple User Crud"
		docs.SwaggerInfo.Description = "This is a sample server to a User CRUD."
		docs.SwaggerInfo.Version = Version
		docs.SwaggerInfo.BasePath = "/" + BasePath
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
		main.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		//Tests
		main.GET("/ping", Ping)
		main.GET("/pong", Pong)

		//User routes
		main.POST("/users", CreateUser)
		users := main.Group("users", core.Auth())
		{
			users.GET("/", ListUsers)
			users.PUT("/", UpdateUser)
			users.DELETE("/delete", DeleteUser)
			users.DELETE("/delete-all", DeleteAll)
		}
	}
}
