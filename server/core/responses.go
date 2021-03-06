package core

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, response gin.H) {
	c.JSON(200, response)
}

func Error(c *gin.Context, errorMessage string, err error) {
	log.Println("log error: ", err)
	c.JSON(400, gin.H{
		"message": errorMessage,
	})
}
func UnknownError(c *gin.Context, err error) {
	log.Println("log fatal error: ", err)
	c.JSON(500, gin.H{
		"message": "Unknown Error",
	})
}
