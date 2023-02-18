package dao

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ParentId uint
	Name     string
	Status   byte
}
