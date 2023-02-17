package dao

import (
	"database/sql"

	"github.com/IvSen/shareThings/pkg/client/postgresql/gorm"
)

type Child struct {
	gorm.Model
	Name     string
	UserID   string
	GenderID sql.NullInt64
}
