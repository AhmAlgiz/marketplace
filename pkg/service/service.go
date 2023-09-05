package service

import (
	"github.com/AhmAlgiz/marketplace/pkg/repository"
	"github.com/AhmAlgiz/marketplace/structures"
)

type Auth interface {
	CreateUser(user structures.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Item interface {
	CreateItem(item structures.Item) (int, error)
	GetAllItems() ([]structures.Item, error)
	GetItemById(id int) ([]structures.Item, error)
	GetItemByTitle(title string) ([]structures.Item, error)
	GetItemByUsername(username string) ([]structures.Item, error)
	DeleteItem(userId, id int) error
	UpdateItem(input structures.UpdateItem, userId int) error
}

type Service struct {
	Auth
	Item
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repos.Auth),
		Item: NewItemService(repos.Item),
	}
}
