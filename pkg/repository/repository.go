package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/zenmaster911/Game/pkg/model"
)

type Authorization interface {
	Create(input *model.User) error
	//GetUser(username, password string) (model.User, error)
	GetByUsername(username string) (model.User, error)
}

type Character interface {
	CreateChar(userId int, char *model.Character) (int, error)
	UserChars(userId int) ([]model.CharacterIntro, error)
	GetByNickname(nickname string) (model.Character, error)
	DeleteCharByNickname(userId int, charNickname string) error
	GetCharById(userId, charId int) (model.Character, error)
	DeleteCharById(userId, charId int) error
}

type Skill interface {
	CreateSkill(skill *model.Skill) error
}
type Repository struct {
	Authorization
	Character
	Skill
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewUserRepository(db),
		Character:     NewCharacterPostgres(db),
	}
}
