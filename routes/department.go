package routes

import (
	"golang-basic-gin/config"
	"golang-basic-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetDepartment(c *gin.Context) {

	var departments []models.Department

	// config.DB.Find(&departments)
	config.DB.Preload("Positions").Find(&departments)

	respDepartment := []models.ResponseGetDepartment{}

	for _, dept := range departments {
		respPost := []models.ResponseEachPosition{}
		for _, post := range dept.Positions {
			resp := models.ResponseEachPosition{
				Id:   post.ID,
				Name: post.Name,
				Code: post.Code,
			}
			respPost = append(respPost, resp)
		}

		respDept := models.ResponseGetDepartment{
			Id:       dept.ID,
			Name:     dept.Name,
			Code:     dept.Code,
			Position: respPost,
		}
		respDepartment = append(respDepartment, respDept)
	}

	c.JSON(200, gin.H{
		"message": "Welcome in department",
		"data":    respDepartment,
	})
}

func GetDepartmentById(c *gin.Context) {
	id := c.Param("id")
	var department models.Department

	data := config.DB.Preload(clause.Associations).First(&department, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data Not Found",
			"status":  "Data Not Found",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    department,
	})
}

func PostDepartment(c *gin.Context) {

	var department models.Department

	err := c.ShouldBindJSON(&department)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "invalid input,Check input",
		})
		return
	} else {
		config.DB.Create(&department)

		c.JSON(200, gin.H{
			"message": "succes",
			"data":    department,
		})
	}

}

func PutDepartment(c *gin.Context) {
	id := c.Param("id")
	department := models.Department{}
	reqDepartment := models.Department{}

	c.BindJSON(&reqDepartment)
	config.DB.Model(&department).Where("id = ?", id).Updates(reqDepartment)

	c.JSON(200, gin.H{
		"message": "updated succes",
		"data":    department,
	})

}

func DeleteDepartment(c *gin.Context) {
	id := c.Param("id")
	department := models.Department{}

	config.DB.Delete(&department, id)

	c.JSON(200, gin.H{
		"message": "deleted succes",
	})

}
