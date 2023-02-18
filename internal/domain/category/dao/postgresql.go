package dao

import (
	"context"

	"gorm.io/gorm"
)

type CategoryDAO struct {
	client *gorm.DB
}

func NewCategoryStorage(client *gorm.DB) *CategoryDAO {
	return &CategoryDAO{
		client: client,
	}
}

func (s *CategoryDAO) One(ctx context.Context, id string) (*Category, error) {
	client := s.client.WithContext(ctx)
	var category = &Category{}
	resultQ := client.First(category, "id = ?", id)

	return category, resultQ.Error
}
func (s *CategoryDAO) All(ctx context.Context) ([]*Category, error) {
	client := s.client.WithContext(ctx)
	var Category []*Category
	resultQ := client.Find(&Category)

	return Category, resultQ.Error
}
