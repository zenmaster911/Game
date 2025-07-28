package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/zenmaster911/Game/pkg/model"
)

type CharacterPostgres struct {
	db *sqlx.DB
}

func NewCharacterPostgres(db *sqlx.DB) *CharacterPostgres {
	return &CharacterPostgres{db: db}
}

func (r *CharacterPostgres) Create(userID int, char *model.Character) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("beginninng DB error: %s", err)
	}

	var id int
	query := `INSERT INTO characters ()`
}
