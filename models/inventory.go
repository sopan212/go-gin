package models

import (
	"time"

	"gorm.io/gorm"
)

type Inventory struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `josn:"description"`
	Archive     Archive
	Employees   []EmployeeInventory `json:"employees"`
}

type EmployeeInventory struct {
	gorm.Model
	EmployeeID  uint      `json:"employee_id"`
	Employee    Employee  `gorm:"foreignKey:EmployeeID;reference:ID"`
	InventoryID uint      `json:"inventory_id"`
	Inventory   Inventory `gorm:"foreignKey:InventoryID;refrence:ID"`
	Description string    `json:"description"`
}

type ReqInventory struct {
	InventoryName        string `json:"inventory_name"`
	InventoryDescription string `json:"inventory_description"`
	ArchiveName          string `json:"archive_name"`
	ArchiveDescription   string `json:"archive_description"`
}

type ResponseGetRrentals struct {
	Id            uint   `json:"id"`
	Description   string `json:"description"`
	EmployeeName  string `json:"employee_name"`
	InventoryName string `json:"inventory_name"`
	CreatedAt     time.Time
}
