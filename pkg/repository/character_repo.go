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

func (r *CharacterPostgres) CreateChar(userID int, char *model.Character) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("beginninng DB error: %s", err)
	}

	var id int
	CharQuery := `INSERT INTO characters (user_id, nickname, class)
	VALUES ($1,$2,$3)
	RETURNING id`

	if err := tx.QueryRow(CharQuery, userID, char.Nickname, char.Class).Scan(&id); err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("create char error: %s", err)
	}

	UsersCharQuery := `INSERT INTO users_characters (user_id, character_id)
	VALUES ($1, $2)`

	_, err = tx.Exec(UsersCharQuery, userID, id)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("create users_chars error: %s", err)
	}

	return id, tx.Commit()

}
