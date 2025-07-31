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

func (r *CharacterPostgres) UserChars(userId int) ([]model.CharacterIntro, error) {

	var chars []model.CharacterIntro

	query := "SELECT c.nickname, c.class, c.level FROM characters c INNER JOIN users_characters uc on c.id = uc.character_id WHERE uc.user_id = $1"

	err := r.db.Select(&chars, query, userId)
	if err != nil {
		return nil, fmt.Errorf("error in selecting chars: %s", err)
	}
	return chars, err

}

func (r *CharacterPostgres) DeleteCharByNickname(userId int, charNickname string) error {

	query := "DELETE FROM characters c WHERE c.nickname = $1 AND c.user_id = $2"

	_, err := r.db.Exec(query, charNickname, userId)

	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (r *CharacterPostgres) GetByNickname(nickname string) (model.Character, error) {
	var char model.Character
	err := r.db.Get(&char, "SELECT * FROM characters WHERE nickname=$1", nickname)
	return char, err
}
