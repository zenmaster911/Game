package model

import "encoding/json"

type SkillType string

const (
	Passive SkillType = "passive"
	Active  SkillType = "active"
)

type Skill struct {
	Id            int             `json:"id" db:"id"`
	Name          string          `json:"name" validate:"required,min=4,max=100" db:"name"`
	Description   string          `json:"description" db:"description"`
	SkillType     SkillType       `json:"skill_type" validate:"oneof=passive active" db:"skill_type"`
	Effect        json.RawMessage `json:"effect" validate:"required" db:"effect"`
	RequiredLevel int             `json:"required_level"  db:"required_level"`
	RequiredClass CharacterClass  `json:"required_class" db:"required_class"`
}
