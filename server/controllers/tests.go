package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ogustavobelo/simple-crud-go/core"
	"github.com/ogustavobelo/simple-crud-go/services"
)

func Ping(c *gin.Context) {
	id := c.Query("id")
	token, err := services.NewJWTService().GenerateToken(id)
	if err != nil {
		core.UnknownError(c, err)
		return
	}
	core.Success(c, gin.H{
		"message": "pong",
		"token":   token,
	})
}

func Pong(c *gin.Context) {
	core.Error(c, "Pinggg2", nil)
}
