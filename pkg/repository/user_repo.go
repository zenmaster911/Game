package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/zenmaster911/Game/pkg/model"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *model.User) error {
	query := `INSERT INTO users (name, username, password_hash, role) 
	VALUES (:name, :username, :password_hash, 'player')
	RETURNING id`

	_, err := r.db.NamedExec(query, user)
	return err
}

func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE username=$1", username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUser(username, password string) (model.User, error) {
	var user model.User
	err := r.db.Get(&user, "SELECT id FROM users WHERE username=$1 AND password_hash=$2", username, password)
	return user, err
}
