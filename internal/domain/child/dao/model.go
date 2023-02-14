package dao

import (
	genderDao "github.com/IvSen/shareThings/internal/domain/gender/dao"
	userDao "github.com/IvSen/shareThings/internal/domain/user/dao"

	"github.com/IvSen/shareThings/pkg/client/postgresql/gorm"
)

type Child struct {
	gorm.Model
	Name     string
	UserID   *uint
	User     *userDao.User
	GenderID *uint
	Gender   *genderDao.Gender
}
