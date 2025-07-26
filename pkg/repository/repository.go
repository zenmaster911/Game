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

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewUserRepository(db),
	}
}
