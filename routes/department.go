package routes

import "github.com/gin-gonic/gin"

func GetDepartment(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome in department",
	})
}
