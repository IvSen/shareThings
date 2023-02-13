package service

import (
	"context"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"time"

	"github.com/IvSen/shareThings/internal/controller/http/v1/dto"

	"github.com/IvSen/shareThings/internal/domain/user/dao"
	"github.com/IvSen/shareThings/internal/domain/user/model"
)

const (
	// TODO: вынести в конфиг
	salt     = "hjqrhjqw1246sdfsdf17ajfhsdfsdfajs"
	tokenTTL = 6 * time.Hour
)

type repository interface {
	One(context.Context, string) (*dao.User, error)
	Create(*context.Context, *dao.User) (*dao.User, error)
	Update(*context.Context, *dao.User) (*dao.User, error)
	Delete(*context.Context, uint) error
	GetByLogin(context.Context, string) (*dao.User, error)
	GetByEmailAndPassword(context.Context, string, string) (*dao.User, error)
	//GenerateToken(*context.Context, dto.UserLoginRequest) (*dao.User, error)
}

type UserService struct {
	repository repository
}

func NewUserService(repository repository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) Create(ctx context.Context, user *dto.CreateUserRequest) (*model.User, error) {

	StorageModel := dao.User{
		Name:     sql.NullString{String: user.Name}.String,
		Email:    sql.NullString{String: user.Email}.String,
		Password: s.GeneratePasswordHash(user.Password),
		Country:  sql.NullString{String: user.Country},
		City:     sql.NullString{String: user.City},
		District: sql.NullString{String: user.District},
		Postcode: sql.NullString{String: user.Postcode},
	}
	create, err := s.repository.Create(&ctx, &StorageModel)
	if err != nil {
		return nil, err
	}

	return model.NewUser(create), nil
}

func (s *UserService) Update(ctx context.Context, user *model.User) (*model.User, error) {
	userFromDB, err := s.One(ctx, user.UUID)
	if err != nil {
		return nil, err
	}

	if len(user.Password) > 0 {
		userFromDB.Password = s.GeneratePasswordHash(user.Password)
	}

	if len(user.Email) > 0 {
		userFromDB.Email = user.Email
	}

	if len(user.City) > 0 {
		userFromDB.City = user.City
	}
	if len(user.District) > 0 {
		userFromDB.District = user.District
	}
	if len(user.Postcode) > 0 {
		userFromDB.Postcode = user.Postcode
	}

	//if len(user.PhotoId) > 0 {
	//	userFromDB.PhotoId = user.PhotoId
	//}

	StorageModel := dao.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: s.GeneratePasswordHash(user.Password),
		Country:  sql.NullString{String: user.Country},
		City:     sql.NullString{String: user.City},
		District: sql.NullString{String: user.District},
		Postcode: sql.NullString{String: user.Postcode},
	}

	updatedUser, err := s.repository.Update(&ctx, &StorageModel)
	if err != nil {
		return nil, err
	}

	return model.NewUser(updatedUser), nil
}

func (s *UserService) One(ctx context.Context, id string) (*model.User, error) {
	one, err := s.repository.One(ctx, id)
	if err != nil {
		return nil, err
	}

	return model.NewUser(one), nil
}
func (s *UserService) GetByLogin(ctx context.Context, email string) (*dao.User, error) {
	one, err := s.repository.GetByLogin(ctx, email)
	if err != nil {
		return nil, err
	}

	return one, nil
	//return model.NewUser(one), nil
}
func (s *UserService) GetByEmailAndPassword(ctx context.Context, email string, password string) (*dao.User, error) {
	one, err := s.repository.GetByEmailAndPassword(ctx, email, password)
	if err != nil {
		return nil, err
	}

	return one, nil
}

func (s *UserService) Delete(ctx *context.Context, id uint) error {
	return s.repository.Delete(ctx, id)
}

func (s *UserService) GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
