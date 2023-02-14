package dao

import (
	"gorm.io/gorm"
)

type Gender struct {
	gorm.Model
	Name string
}
