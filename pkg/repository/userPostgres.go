package repository

import (
	"fmt"

	"github.com/AhmAlgiz/marketplace/structures"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) UpdateUser(updateUser structures.UpdateUser, userId int) error {
	query := fmt.Sprintf("UPDATE %s ut SET username=$1 WHERE ut.id=$2", usersTable)

	_, err := r.db.Exec(query, updateUser.Name, userId)

	return err
}
