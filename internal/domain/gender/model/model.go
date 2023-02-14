package model

import (
	"github.com/IvSen/shareThings/internal/domain/gender/dao"
)

type Gender struct {
	Id   uint
	Name string
}

func DtoToDb(sp *dao.Gender) *Gender {
	return &Gender{
		Id:   sp.Model.ID,
		Name: sp.Name,
	}
}
