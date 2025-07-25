package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/zenmaster911/Game/pkg/model"
)

type UerRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UerRepository {
	return &UerRepository{db: db}
}

func (r *UerRepository) Create(user *model.User) error {
	query := `INSERT INTO users (name, username, password_hash, role) 
	VALUES (:name, :username, :password 'player')
	RETURNING id`

	_, err := r.db.NamedExec(query, user)
	return err
}

func (r *UerRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE username = $1", username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
