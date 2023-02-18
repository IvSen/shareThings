package model

import (
	"github.com/IvSen/shareThings/internal/domain/category/dao"
)

type Category struct {
	Id       uint
	Name     string
	ParentId uint
	Status   byte
}

func DtoToDb(sp *dao.Category) *Category {
	return &Category{
		Id:       sp.Model.ID,
		Name:     sp.Name,
		ParentId: sp.ParentId,
		Status:   sp.Status,
	}
}
