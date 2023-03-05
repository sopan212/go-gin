package routes

import (
	"golang-basic-gin/config"
	"golang-basic-gin/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func PostEmploye(c *gin.Context) {
	var employes models.Employee

	c.ShouldBindJSON(&employes)

	// if err != nil {
	// 	c.AbortWithStatusJSON(400, gin.H{
	// 		"err":     err.Error(),
	// 		"message": "input invalid check input!!",
	// 	})
	// 	c.Abort()
	// 	return
	// }
	// err = config.DB.Where("phone=?,email=?", employes.Phone, employes.Email).First(&employes).Error
	// if err == nil {
	// 	c.JSON(409, gin.H{
	// 		"message": "no telp atau email sudah terdaftar",
	// 	})
	// 	return
	// }
	config.DB.Create(&employes)
	c.JSON(200, gin.H{
		"message": "success",
		"data":    employes,
	})
}

func GetEmployee(c *gin.Context) {
	var employes []models.Employee

	config.DB.Preload(clause.Associations).Find(&employes)

	c.JSON(200, gin.H{
		"message": "success",
		"data":    employes,
	})

}

func PutEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee
	var reqEmployee models.Employee

	c.BindJSON(&reqEmployee)
	config.DB.Model(&employee).Where("id = ?", id).Updates(reqEmployee)
	c.JSON(200, gin.H{
		"message": "update success",
		"data":    employee,
	})

}

func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	employee := models.Employee{}

	config.DB.Delete(&employee, id)

	c.JSON(200, gin.H{
		"message": "deleted succes",
	})

}
