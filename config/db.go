package config

import (
	"golang-basic-gin/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/gin_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.Department{}, models.Position{}, &models.Employee{}, &models.Archive{}, &models.Inventory{})

	// DB.Create(&models.Articels{
	// Title:"Naruto Shipuden",
	// Author:"Masasi Kisimoto",
	// })

	// DB.Create(&models.Department{
	// 	Name: "Finance",
	// 	Code: "Fin",
	// 	Positions: []models.Position{
	// 		{Name: "General Manager Finance", Code: "GMF"},
	// 		{Name: "Manager Finance", Code: "MF"},
	// 		{Name: "Staff Finance", Code: "SF"},
	// 	},
	// })

	// DB.Create(&models.Department{
	// 	Name: "Information & Technology",
	// 	Code: "IT",
	// 	Positions: []models.Position{
	// 		{Name: "General Manager IT", Code: "GMI"},
	// 		{Name: "Engginer Manager", Code: "EM"},
	// 		{Name: "Engginer", Code: "EN"},
	// 	},
	// })
}
