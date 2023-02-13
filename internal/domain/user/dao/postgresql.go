package dao

import (
	"context"
	"errors"
	"fmt"

	gormModel "github.com/IvSen/shareThings/pkg/client/postgresql/gorm"
	"gorm.io/gorm"
)

type UserDAO struct {
	client *gorm.DB
}

func NewUserStorage(client *gorm.DB) *UserDAO {
	return &UserDAO{
		client: client,
	}
}

func (s *UserDAO) Create(ctx *context.Context, user *User) (*User, error) {
	client := s.client.WithContext(*ctx)

	var existUser = &User{}
	firstResult := client.Where("email = ?", user.Email).First(existUser)

	if firstResult.Error == nil {
		// TODO: LOG
		fmt.Println("User already exist")
		return nil, errors.New("user already exist")
	}

	result := client.Create(&user)
	if result.Error != nil {
		// TODO: log
		fmt.Println(result.Error)
	}

	return user, result.Error
}

func (s *UserDAO) One(ctx context.Context, id string) (*User, error) {
	client := s.client.WithContext(ctx)
	var user = &User{}
	resultQ := client.Model(User{
		Model: gormModel.Model{Id: id},
	}).First(&user)

	return user, resultQ.Error
}

func (s *UserDAO) GetByEmailAndPassword(ctx context.Context, email string, password string) (*User, error) {
	client := s.client.WithContext(ctx)
	var user = &User{}
	resultQ := client.Model(User{
		Email:    email,
		Password: password,
	}).First(&user)

	return user, resultQ.Error
}

func (s *UserDAO) Update(ctx *context.Context, user *User) (*User, error) {
	client := s.client.WithContext(*ctx)
	result := client.Save(&user)
	if result.Error != nil {
		// TODO: log
		fmt.Println(result.Error)
	}
	return user, result.Error
}
