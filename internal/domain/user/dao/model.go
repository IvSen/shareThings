package dao

import (
	"database/sql"

	"github.com/IvSen/shareThings/pkg/client/postgresql/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Country  sql.NullString
	City     sql.NullString
	District sql.NullString
	Postcode sql.NullString
}
