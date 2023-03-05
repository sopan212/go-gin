package models

import "gorm.io/gorm"

type Archive struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `josn:"description"`
	inventoryID uint
}
