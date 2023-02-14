package dao

import (
	"context"

	"gorm.io/gorm"
)

type GenderDAO struct {
	client *gorm.DB
}

func NewGenderStorage(client *gorm.DB) *GenderDAO {
	return &GenderDAO{
		client: client,
	}
}

func (s *GenderDAO) One(ctx context.Context, id uint) (*Gender, error) {
	client := s.client.WithContext(ctx)
	var gender = &Gender{}
	resultQ := client.Model(Gender{
		Model: gorm.Model{ID: id},
	}).First(&gender)

	return gender, resultQ.Error
}
func (s *GenderDAO) All(ctx context.Context) ([]*Gender, error) {
	client := s.client.WithContext(ctx)
	var gender []*Gender
	resultQ := client.Find(&gender)

	return gender, resultQ.Error
}
