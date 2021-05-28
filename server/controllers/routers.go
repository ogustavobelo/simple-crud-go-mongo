package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ogustavobelo/simple-crud-go/core"
)

func InitRoutes(router *gin.Engine) {
	main := router.Group("simple-crud")
	{
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
