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
}

type Repository struct {
	Authorization
	Character
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewUserRepository(db),
		Character:     NewCharacterPostgres(db),
	}
}
