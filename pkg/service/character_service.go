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
func (s *CharacterService) UserChars(userId int) ([]model.CharacterIntro, error) {
	return s.repo.UserChars(userId)
}

func (s *CharacterService) DeleteCharByNickname(userId int, charNickname string) error {
	return s.repo.DeleteCharByNickname(userId, charNickname)
}

func (s *CharacterService) GetByNickname(nickname string) (model.Character, error) {
	return s.repo.GetByNickname(nickname)
}
