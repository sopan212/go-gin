package routes

import "github.com/gin-gonic/gin"

func GetPosition(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome in Position",
	})
}
