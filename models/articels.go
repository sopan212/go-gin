package models

import (
	"gorm.io/gorm"
)

type Articels struct {
	gorm.Model
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// type Articels []Articel
