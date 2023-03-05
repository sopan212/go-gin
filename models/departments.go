package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Name      string     `json:"name" binding:"required,min=1"`
	Code      string     `json:"code" binding:"required"`
	Positions []Position `json:"positions"`
}
