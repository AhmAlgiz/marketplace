package repository

import "github.com/jmoiron/sqlx"

type Auth interface {
}

type Item interface {
}

type Repository struct {
	Auth
	Item
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
