package service

import (
	"github.com/zenmaster911/Game/pkg/model"
	"github.com/zenmaster911/Game/pkg/repository"
)

type CharacterService struct {
	repo repository.Character
}

func NewCharacterService(repo repository.Character) *CharacterService {
	return &CharacterService{repo: repo}
}

func (s *CharacterService) CreateChar(userID int, char *model.Character) (int, error) {
	return s.repo.CreateChar(userID, char)
}
