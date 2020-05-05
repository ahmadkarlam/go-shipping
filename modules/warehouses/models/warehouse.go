package models

import "github.com/jinzhu/gorm"

type Warehouse struct {
	gorm.Model
	Code  string `gorm:"column:code"`
	Stock int    `gorm:"column:stock"`
	X     int    `gorm:"column:x"`
	Y     int    `gorm:"column:y"`
}
