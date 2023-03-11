package models

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name       string `json:"name" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
	Address    string `json:"address" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	PositionID uint   `json:"position_id"`
	Position   Position
	//Inventories []Inventory `gorm:"many2many:employee_inventories;"`
	Inventories []EmployeeInventory `json:"inventories"`
}
