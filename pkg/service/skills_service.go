package service

import (
	"github.com/zenmaster911/Game/pkg/model"
	"github.com/zenmaster911/Game/pkg/repository"
)

type SkillService struct {
	repo repository.Skill
}

func NewSkillService(repo repository.Skill) *SkillService {
	return &SkillService{repo: repo}
}

func (s *SkillService) CreateSkill(skill *model.Skill) error {
	return s.repo.CreateSkill(skill)
}
