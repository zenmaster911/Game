package service

import (
	"errors"

	"github.com/zenmaster911/Game/pkg/model"
	"github.com/zenmaster911/Game/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UerRepository
}

func NewUserService(repo *repository.UerRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) SignUp(input *model.CreateUser) (*model.User, error) {
	_, err := s.repo.GetByUsername(input.Username)
	if err == nil {
		return nil, errors.New("user with this name already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:     input.Name,
		Username: input.Username,
		Password: string(hashedPassword),
		Role:     "player",
	}

	err = s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	user.Password = ""

	return user, nil
}
