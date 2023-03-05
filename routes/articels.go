package routes

import (
	"golang-basic-gin/config"
	"golang-basic-gin/models"

	"github.com/gin-gonic/gin"
)

func GetArticles(c *gin.Context) {

	articles := []models.Articels{}
	config.DB.Find(&articles)
	c.JSON(200, gin.H{
		"message": "Articless Page",
		"data":    articles,
	})
}

func PostArticles(c *gin.Context) {

	var article models.Articels
	// articel := models.Articels{
	// 	Title: c.PostForm("title"),
	// 	Author:c.PostForm("author"),
	// }

	err := c.ShouldBindJSON(&article)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"err":     err.Error(),
			"message": "Input required fill",
		})
		return
	}
	config.DB.Create(&article)

	c.JSON(200, gin.H{
		"message": "success",
		"data":    article,
	})
}

func Putarticle(c *gin.Context) {
	id := c.Param("id")

	var article models.Articels

	var reqArtilces models.Articels

	c.BindJSON(&reqArtilces)

	config.DB.Model(&article).Where("id = ?", id).Updates(reqArtilces)

	c.JSON(200, gin.H{
		"message": "updated",
		"data":    article,
	})
}

func DeleteArticle(c *gin.Context) {
	id := c.Param("id")

	var articel models.Articels

	config.DB.Delete(&articel, id)

	c.JSON(200, gin.H{
		"message": "Delleted",
	})
}
