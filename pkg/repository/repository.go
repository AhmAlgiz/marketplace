package repository

import (
	"github.com/AhmAlgiz/marketplace/structures"
	"github.com/jmoiron/sqlx"
)

type Auth interface {
	CreateUser(user structures.User) (int, error)
	GetUserByName(username string) (structures.User, error)
}

type Item interface {
	CreateItem(input structures.Item) (int, error)
}

type Repository struct {
	Auth
	Item
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthPostgres(db),
		Item: NewItemPostgres(db),
	}
}
