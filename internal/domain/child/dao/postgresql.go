package dao

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type ChildDAO struct {
	client *gorm.DB
}

func NewGenderStorage(client *gorm.DB) *ChildDAO {
	return &ChildDAO{
		client: client,
	}
}

func (s *ChildDAO) Delete(ctx context.Context, id string) error {
	client := s.client.WithContext(ctx)
	result := client.Delete(&Child{}, "id = ?", id)
	return result.Error
}

func (s *ChildDAO) One(ctx context.Context, id string) (*Child, error) {
	var userId = ctx.Value("user_id")
	client := s.client.WithContext(ctx)
	var user = &Child{}
	resultQ := client.First(user, "id = ?", id).Where("user_id = ?", userId)

	return user, resultQ.Error
}

func (s *ChildDAO) All(ctx context.Context) ([]*Child, error) {
	var userId = ctx.Value("user_id")

	client := s.client.WithContext(ctx)
	var child []*Child
	resultQ := client.Find(&child, "user_id = ?", userId)

	return child, resultQ.Error
}

func (s *ChildDAO) Create(ctx context.Context, child *Child) (*Child, error) {
	client := s.client.WithContext(ctx)

	result := client.Create(&child)
	if result.Error != nil {
		// TODO: log
		//fmt.Println(result.Error)
	}

	return child, result.Error
}

func (s *ChildDAO) Update(ctx context.Context, child *Child) (*Child, error) {
	client := s.client.WithContext(ctx)
	result := client.Save(&child)
	if result.Error != nil {
		// TODO: log
		fmt.Println(result.Error)
	}
	return child, result.Error
}
