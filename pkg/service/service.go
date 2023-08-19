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
