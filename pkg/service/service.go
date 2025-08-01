package service

import (
	"github.com/zenmaster911/Game/pkg/model"
	"github.com/zenmaster911/Game/pkg/repository"
)

type Character interface {
	CreateChar(userId int, char *model.Character) (int, error)
	UserChars(userId int) ([]model.CharacterIntro, error)
	DeleteCharByNickname(userId int, charNickname string) error
	GetByNickname(nickname string) (model.Character, error)
	GetCharById(userId, charId int) (model.Character, error)
}

type Authorization interface {
	Create(input *model.CreateUser) (*model.User, error)
	ParseToken(accessToken string) (int, error)
	GenerateToken(username, password string) (string, error)
}
type Service struct {
	Authorization
	Character
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewUserService(repo.Authorization),
		Character:     NewCharacterService(repo.Character),
	}
}
