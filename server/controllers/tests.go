package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ogustavobelo/simple-crud-go/core"
)

func Ping(c *gin.Context) {
	core.Success(c, gin.H{"message": "pong"})
}

func Pong(c *gin.Context) {
	core.Error(c, "Pinggg2", nil)
}