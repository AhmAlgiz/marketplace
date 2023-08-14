package repository

import (
	"fmt"

	"github.com/AhmAlgiz/marketplace/structures"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user structures.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, pass_hash) VALUES ($1, $2) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Pass)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
