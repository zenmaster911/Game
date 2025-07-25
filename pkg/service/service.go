package service

import (
	"github.com/zenmaster911/Game/pkg/model"
	"github.com/zenmaster911/Game/pkg/repository"
)

type Authorization interface {
	Create(input *model.CreateUser) (*model.User, error)
}
type Service struct {
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewUserService(repo.Authorization),
	}
}
