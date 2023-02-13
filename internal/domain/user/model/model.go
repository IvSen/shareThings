package model

import (
	"github.com/IvSen/shareThings/internal/domain/user/dao"
)

type User struct {
	UUID     string
	Name     string
	Email    string
	Password string
	Country  string
	City     string
	District string
	Postcode string
	//PhotoId   *int64
}

func NewUser(sp *dao.User) *User {
	return &User{
		UUID:     sp.Model.UUID,
		Name:     sp.Name,
		Email:    sp.Email,
		Country:  sp.Country.String,
		City:     sp.City.String,
		District: sp.District.String,
		Postcode: sp.Postcode.String,
	}
}
