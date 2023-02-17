package model

import "github.com/IvSen/shareThings/internal/domain/child/dao"

type Child struct {
	Id       string
	Name     string
	UserID   string
	GenderID int64
}

func DtoToDb(sp *dao.Child) *Child {
	return &Child{
		Id:       sp.Model.Id,
		Name:     sp.Name,
		UserID:   sp.UserID,
		GenderID: sp.GenderID.Int64,
	}
}
