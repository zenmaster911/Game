package model

import "encoding/json"

type SkillClass string

type Skill struct {
	Id            int             `json:"id" db:"id"`
	Name          string          `json:"name" validate:"required, min=4,max=100" db:"name"`
	Description   string          `json:"description" db:"description"`
	SkillType     string          `json:"skill_type" validate:"oneof=passive active" db:"skill_type"`
	Effect        json.RawMessage `json:"effect" validate:"required" db:"effect"`
	RequiredLevel int             `json:"required_level"  db:"required_level"`
	RequiredClass CharacterClass  `json:"character_class" db:"character_class"`
}
