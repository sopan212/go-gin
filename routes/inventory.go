package routes

import (
	"golang-basic-gin/config"
	"golang-basic-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func PostInventory(c *gin.Context) {
	var reqInventory models.ReqInventory

	c.BindJSON(&reqInventory)

	inventory := models.Inventory{
		Name:        reqInventory.InventoryName,
		Description: reqInventory.InventoryDescription,
		Archive: models.Archive{
			Name:        reqInventory.ArchiveName,
			Description: reqInventory.ArchiveDescription,
		},
	}

	data := config.DB.Create(&inventory)
	if data.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "failed to save data",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    inventory,
	})
}

func GetInventory(c *gin.Context) {
	var inventory []models.Inventory

	config.DB.Preload(clause.Associations).Find(&inventory)

	c.JSON(200, gin.H{
		"message": "success",
		"data":    inventory,
	})

}
func GetInventoryById(c *gin.Context) {

	id := c.Param("id")
	var inventory models.Inventory

	config.DB.Preload(clause.Associations).First(&inventory, "id = ?", id)

	c.JSON(200, gin.H{
		"message": "success",
		"data":    inventory,
	})

}

func PutInventory(c *gin.Context) {
	id := c.Param("id")

	var reqInventory models.ReqInventory

	c.BindJSON(&reqInventory)
	var inventory models.Inventory
	updateInventory := models.Inventory{
		Name:        reqInventory.InventoryName,
		Description: reqInventory.InventoryDescription,
	}

	data := config.DB.Model(&inventory).Where("id = ?", id).Updates(updateInventory)

	if data.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"messagae": "failed updated data",
			"staus":    "bad request",
		})
		return
	}

	var archive models.Archive

	updateArchive := models.Archive{
		Name:        reqInventory.ArchiveName,
		Description: reqInventory.ArchiveDescription,
	}

	data = config.DB.Model(&archive).Where("inventory_id = ?", id).Updates(updateArchive)
	if data.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "failed update",
			"status":  "bad request",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "update success",
		"data":    inventory,
	})

}

func DeleteInventory(c *gin.Context) {
	id := c.Param("id")
	inventory := models.Inventory{}

	config.DB.Delete(&inventory, id)

	c.JSON(200, gin.H{
		"message": "deleted succes",
	})

}
