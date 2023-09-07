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

func (r *UserPostgres) GetUserById(id int) ([]structures.GetUser, error) {
	query := fmt.Sprintf(
		`SELECT id, username FROM %s WHERE id=$1`, usersTable)

	var sl []structures.GetUser

	err := r.db.Select(&sl, query, id)

	return sl, err
}
