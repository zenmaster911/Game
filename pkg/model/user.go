package model

type User struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Username string `json:"username" db:"username"`
	Password string `json:"password_hash" db:"password_hash"`
	Role     string `json:"role" db:"role"`
}

type CreateUser struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Username string `json:"username" validate:"required,min=4,max=100"`
	Password string `json:"password" validate:"required,min=6,max=255"`
}
