package main

import (
	"golang-basic-gin/config"
	"golang-basic-gin/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	route := gin.Default()
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := route.Group("api/v1")
	{
		v1.GET("/", GetHome)

		department := v1.Group("/department")
		{
			department.GET("/", routes.GetDepartment)
			department.GET("/:id", routes.GetDepartmentById)
			department.POST("/", routes.PostDepartment)
			department.PUT("/:id", routes.PutDepartment)
			department.DELETE("/:id", routes.DeleteDepartment)
		}
		employee := v1.Group("/employee")
		{
			employee.POST("/", routes.PostEmploye)
			employee.GET("/", routes.GetEmployee)
			employee.PUT("/:id", routes.PutEmployee)
			employee.DELETE("/:id", routes.DeleteEmployee)
		}
		articles := v1.Group("/articles")
		{
			articles.GET("/", routes.GetArticles)
			articles.POST("/", routes.PostArticles)
			articles.PUT("/:id", routes.Putarticle)
			articles.DELETE("/:id", routes.DeleteArticle)
		}
		position := v1.Group("/position")
		{
			position.GET("/", routes.GetPosition)
			position.POST("/", routes.PostPosition)
			position.PUT("/:id", routes.PutPosition)
			position.DELETE("/:id", routes.DeletePosition)
		}
		inventory := v1.Group("/inventory")
		{
			inventory.GET("/", routes.GetInventory)
			inventory.GET("/:id", routes.GetInventoryById)
			inventory.POST("/", routes.PostInventory)
			inventory.PUT("/:id", routes.PutInventory)
			inventory.DELETE("/:id", routes.DeleteInventory)
		}
	}

	route.Run(":3030") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func GetHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome Home",
	})
}
