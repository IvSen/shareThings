package gorm

import (
	"database/sql"
	"time"
)

type Model struct {
	UUID      string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
