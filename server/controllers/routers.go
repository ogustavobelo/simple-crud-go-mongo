package controllers

import "github.com/gin-gonic/gin"

func InitRoutes(router *gin.Engine) {
	//Tests
	router.GET("/ping", Ping)
	router.GET("/pong", Pong)

	//Users Routes
	router.POST("/users/create", CreateUser)
	router.GET("/users", ListUsers)
	router.PUT("/users/update", UpdateUser)
	router.DELETE("/users/delete-all", DeleteAll)
}
