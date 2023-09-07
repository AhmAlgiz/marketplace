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
	GetAllItems() ([]structures.Item, error)
	GetItemById(id int) ([]structures.Item, error)
	GetItemByTitle(title string) ([]structures.Item, error)
	GetItemByUsername(username string) ([]structures.Item, error)
	DeleteItem(id, userId int) error
	UpdateItem(input structures.UpdateItem, userId int) error
}

type User interface {
	UpdateUser(updateUser structures.UpdateUser, userId int) error
	GetUserById(id int) ([]structures.GetUser, error)
}

type Repository struct {
	Auth
	Item
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthPostgres(db),
		Item: NewItemPostgres(db),
		User: NewUserPostgres(db),
	}
}
