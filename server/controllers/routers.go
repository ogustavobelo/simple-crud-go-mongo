package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ogustavobelo/simple-crud-go/core"
	"github.com/ogustavobelo/simple-crud-go/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(router *gin.Engine) {
	main := router.Group("simple-crud")
	{
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
		main.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		//Tests
		main.GET("/ping", Ping)
		main.GET("/pong", Pong)

		//Users Routes
		main.POST("/users", CreateUser)
		users := main.Group("users", core.Auth())
		{
			users.GET("/", ListUsers)
			users.PUT("/", UpdateUser)
			users.DELETE("/delete-all", DeleteAll)
		}
	}
}
