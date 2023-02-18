package service

import (
	"context"

	"github.com/IvSen/shareThings/internal/domain/category/dao"
	"github.com/IvSen/shareThings/internal/domain/category/model"
)

type repository interface {
	One(context.Context, string) (*dao.Category, error)
	All(context.Context) ([]*dao.Category, error)
}

type CategoryService struct {
	repository repository
}

func NewCategoryService(repository repository) *CategoryService {
	return &CategoryService{repository: repository}
}

func (s *CategoryService) One(ctx context.Context, id string) (*model.Category, error) {
	one, err := s.repository.One(ctx, id)
	if err != nil {
		return nil, err
	}

	return model.DtoToDb(one), nil
}

func (s *CategoryService) All(ctx context.Context) ([]*model.Category, error) {
	all, err := s.repository.All(ctx)
	if err != nil {
		return nil, err
	}
	allCategory := make([]*model.Category, 0)
	for _, Category := range all {
		allCategory = append(allCategory, model.DtoToDb(Category))
	}

	return allCategory, nil
}
