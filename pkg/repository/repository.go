package repository

import (
	"github.com/AhmAlgiz/marketplace/structures"
	"github.com/jmoiron/sqlx"
)

type Auth interface {
	CreateUser(user structures.User) (int, error)
}

type Item interface {
}

type Repository struct {
	Auth
	Item
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthPostgres(db),
	}
}
