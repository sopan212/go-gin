package models

import "gorm.io/gorm"

type Position struct {
	gorm.Model
	Name         string `json:"name"`
	Code         string `json:"code"`
	DepartmentID uint   `json:"department_id"`
	Department   Department
	Employees    []Employee `json:"employee"`
}

type ResponseEachPosition struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}
