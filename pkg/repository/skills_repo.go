package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/zenmaster911/Game/pkg/model"
)

type SkillsPostgres struct {
	db *sqlx.DB
}

func NewSkillPostgres(db *sqlx.DB) *SkillsPostgres {
	return &SkillsPostgres{db: db}
}

func (r *SkillsPostgres) CreateSkill(skill *model.Skill) error {
	query := `INSERT INTO skills (name, description,skill_type,efect,required_level,required_class)
	VALUES ( :name, :description,:skill_type,:efect,:required_level,:required_class)
	RETURNING id`
	_, err := r.db.NamedExec(query, skill)
	if err != nil {
		return fmt.Errorf("error in creating skill: %s", err)
	}
	return nil
}
