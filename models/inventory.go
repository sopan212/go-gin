package models

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `josn:"description"`
	Archive     Archive
}

type ReqInventory struct {
	InventoryName        string `json:"inventory_name"`
	InventoryDescription string `json:"inventory_description"`
	ArchiveName          string `json:"archive_name"`
	ArchiveDescription   string `json:"archive_description"`
}
