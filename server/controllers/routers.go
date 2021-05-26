package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ogustavobelo/simple-crud-go/core"
)

func InitRoutes(router *gin.Engine) {
	//Tests
	router.GET("/ping", Ping)
	router.GET("/pong", Pong)

	//Users Routes
	router.POST("/users", CreateUser)
	users := router.Group("users", core.Auth())
	{
		users.GET("/", ListUsers)
		users.PUT("/", UpdateUser)
		users.DELETE("/delete-all", DeleteAll)
	}
}
