package routes

import (
	"golang-basic-gin/config"
	"golang-basic-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type Reqrental struct {
	EmployeeID  uint
	InventoryID uint
	Description string
}

func RentalByEmployee(c *gin.Context) {
	var reqRental Reqrental

	if err := c.ShouldBindJSON(&reqRental); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
			"err":     err.Error(),
		})
		c.Abort()
		return
	}

	rental := models.EmployeeInventory{
		EmployeeID:  reqRental.EmployeeID,
		InventoryID: reqRental.InventoryID,
		Description: reqRental.Description,
	}

	insert := config.DB.Create(&rental)
	if insert.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server errors",
			"error":   insert.Error.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "insert succes",
		"data":    rental,
	})
}

func GetRentalByInventoryID(c *gin.Context) {
	id := c.Param("id")
	var Inventories []models.Inventory

	config.DB.Preload("Employees").First(&Inventories, "id = ?", id)

	c.JSON(200, gin.H{
		"message": "welcome to data rental",
		"data":    Inventories,
	})
}

func GetRental(c *gin.Context) {
	EmployeeInventory := []models.EmployeeInventory{}

	config.DB.Preload(clause.Associations).Find(&EmployeeInventory)

	responseGetRrentals := []models.ResponseGetRrentals{}
	for _, empInve := range EmployeeInventory {
		respGetRent := models.ResponseGetRrentals{
			Id:            empInve.ID,
			Description:   empInve.Description,
			EmployeeName:  empInve.Employee.Name,
			InventoryName: empInve.Inventory.Name,
			CreatedAt:     empInve.CreatedAt,
		}

		responseGetRrentals = append(responseGetRrentals, respGetRent)
	}

	c.JSON(http.StatusOK, gin.H{
		"messagae": "welcome to data rental",
		"data":     EmployeeInventory,
	})
}
