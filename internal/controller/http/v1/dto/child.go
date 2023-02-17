package dto

import "time"

type CreateUpdateChildRequest struct {
	Id       string
	Name     string
	BirthAT  time.Time
	GenderId int64
	height   float32
}
