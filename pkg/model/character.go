package model

type Character struct {
	ID           int    `json:"id" db:"id"`
	UserID       int    `json:"user_id" db:"user_id"`
	Nickname     string `json:"nickname" db:"nickname"`
	Class        string `json:"class" db:"class"`
	Level        int    `json:"lvl" db:"lvl"`
	Exp          int    `json:"exp" db:"exp"`
	Health       int    `json:"health" db:"health"`
	Strength     int    `json:"strength" db:"strength"`
	Agility      int    `json:"agility" db:"agility"`
	Charisma     int    `json:"charisma" db:"agility"`
	Intelligence int    `json:"intelligence" db:"agility"`
	CreatedTime  string `json:"created_time" db:"created_time"`
}
