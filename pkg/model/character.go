package model

import "time"

type CharacterClass string

const (
	Warrior CharacterClass = "warrior"
	Mage    CharacterClass = "mage"
)

//var validClasses = []CharacterClass{Mage,Warrior}

type Character struct {
	ID           int            `json:"id" db:"id"`
	UserID       int            `json:"user_id" db:"user_id"`
	Nickname     string         `json:"nickname" validate:"required,min=4,max=100" db:"nickname" `
	Class        CharacterClass `json:"class" validate:"oneof=warrior mage" db:"class"`
	Level        int            `json:"lvl" db:"level"`
	Exp          int            `json:"exp" db:"exp"`
	Health       int            `json:"health" db:"health"`
	Strength     int            `json:"strength" db:"strength"`
	Agility      int            `json:"agility" db:"agility"`
	Charisma     int            `json:"charisma" db:"charisma"`
	Intelligence int            `json:"intelligence" db:"intelligence"`
	CreatedTime  time.Time      `json:"created_time" db:"created_time"`
}

type CharacterIntro struct {
	Nickname string         `json:"nickname" db:"nickname" `
	Class    CharacterClass `json:"class" validate:"oneof=warrior mage" db:"class"`
	Level    int            `json:"lvl" db:"level"`
}
