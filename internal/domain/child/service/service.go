package service

import (
	"context"
	"database/sql"

	"github.com/IvSen/shareThings/internal/controller/http/v1/dto"

	"github.com/IvSen/shareThings/internal/domain/child/dao"
	"github.com/IvSen/shareThings/internal/domain/child/model"
)

type repository interface {
	One(context.Context, string) (*dao.Child, error)
	All(context.Context) ([]*dao.Child, error)
	Create(context.Context, *dao.Child) (*dao.Child, error)
	Update(context.Context, *dao.Child) (*dao.Child, error)
	Delete(context.Context, string) error
}

type ChildService struct {
	repository repository
}

func NewChildService(repository repository) *ChildService {
	return &ChildService{repository: repository}
}

func (s *ChildService) Create(ctx context.Context, dto *dto.CreateUpdateChildRequest) (*model.Child, error) {
	StorageModel := dao.Child{
		Name:     dto.Name,
		UserID:   ctx.Value("user_id").(string),
		GenderID: sql.NullInt64{Int64: dto.GenderId},
	}
	create, err := s.repository.Create(ctx, &StorageModel)
	if err != nil {
		return nil, err
	}

	return model.DtoToDb(create), nil
}

func (s *ChildService) Update(ctx context.Context, dto *dto.CreateUpdateChildRequest) (*model.Child, error) {
	userFromDB, err := s.One(ctx, dto.Id)
	if err != nil {
		return nil, err
	}

	if len(dto.Name) > 0 {
		userFromDB.Name = dto.Name
	}
	if dto.GenderId > 0 {
		userFromDB.GenderID = dto.GenderId
	}

	StorageModel := dao.Child{
		Name:     userFromDB.Name,
		UserID:   ctx.Value("user_id").(string),
		GenderID: sql.NullInt64{Int64: 0},
	}

	updatedUser, err := s.repository.Update(ctx, &StorageModel)
	if err != nil {
		return nil, err
	}

	return model.DtoToDb(updatedUser), nil
}

func (s *ChildService) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}

func (s *ChildService) One(ctx context.Context, id string) (*model.Child, error) {
	one, err := s.repository.One(ctx, id)
	if err != nil {
		return nil, err
	}

	return model.DtoToDb(one), nil
}

func (s *ChildService) All(ctx context.Context) ([]*model.Child, error) {
	all, err := s.repository.All(ctx)
	if err != nil {
		return nil, err
	}
	allChild := make([]*model.Child, 0)
	for _, gender := range all {
		allChild = append(allChild, model.DtoToDb(gender))
	}

	return allChild, nil
}
