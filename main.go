package main

import (
	"golang-basic-gin/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := route.Group("api/v1/")
	{
		v1.GET("/", GetHome)

		department := route.Group("/department")
		{
			department.GET("/", routes.GetDepartment)
		}

		position := route.Group("/position")
		{
			position.GET("/position", routes.GetPosition)
		}
	}

	route.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func GetHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome Home",
	})
}
