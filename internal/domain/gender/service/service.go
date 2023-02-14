package service

import (
	"context"

	"github.com/IvSen/shareThings/internal/domain/gender/model"

	"github.com/IvSen/shareThings/internal/domain/gender/dao"
)

type repository interface {
	One(context.Context, uint) (*dao.Gender, error)
	All(context.Context) ([]*dao.Gender, error)
}

type GenderService struct {
	repository repository
}

func NewGenderService(repository repository) *GenderService {
	return &GenderService{repository: repository}
}

func (s *GenderService) One(ctx context.Context, id uint) (*model.Gender, error) {
	one, err := s.repository.One(ctx, id)
	if err != nil {
		return nil, err
	}

	return model.DtoToDb(one), nil
}
func (s *GenderService) All(ctx context.Context) ([]*model.Gender, error) {
	all, err := s.repository.All(ctx)
	if err != nil {
		return nil, err
	}
	allGender := make([]*model.Gender, 0)
	for _, gender := range all {
		allGender = append(allGender, model.DtoToDb(gender))
	}

	return allGender, nil
}
