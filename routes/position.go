package routes

import (
	"golang-basic-gin/config"
	"golang-basic-gin/models"

	"github.com/gin-gonic/gin"
)

func GetPosition(c *gin.Context) {

	position := []models.Position{}
	// config.DB.Find(&position)

	config.DB.Preload("Department").Find(&position)

	respPostion := []models.ResponseEachPosition{}

	for _, post := range position {
		resp := models.ResponseEachPosition{
			Id:   post.ID,
			Name: post.Name,
			Code: post.Code,
		}

		respPostion = append(respPostion, resp)
	}
	c.JSON(200, gin.H{
		"message": "Position Page",
		"data":    respPostion,
	})
}

func PostPosition(c *gin.Context) {

	var position models.Position
	// articel := models.Articels{
	// 	Title: c.PostForm("title"),
	// 	Author:c.PostForm("author"),
	// }
	c.BindJSON(&position)
	// err := c.ShouldBindJSON(&position)
	// if err != nil {
	// 	c.AbortWithStatusJSON(400, gin.H{
	// 		"err":     err.Error(),
	// 		"message": "Input required fill",
	// 	})
	// 	return
	// }
	config.DB.Create(&position)

	c.JSON(200, gin.H{
		"message": "success",
		"data":    position,
	})
}

func PutPosition(c *gin.Context) {
	id := c.Param("id")

	var position models.Position

	var reqPosition models.Position

	c.BindJSON(&reqPosition)

	config.DB.Model(&position).Where("id = ?", id).Updates(reqPosition)

	c.JSON(200, gin.H{
		"message": "updated",
		"data":    position,
	})
}

func DeletePosition(c *gin.Context) {
	id := c.Param("id")

	var Position models.Position

	config.DB.Delete(&Position, id)

	c.JSON(200, gin.H{
		"message": "Delleted",
	})
}
