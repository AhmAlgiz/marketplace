package service

import (
	"github.com/AhmAlgiz/marketplace/pkg/repository"
	"github.com/AhmAlgiz/marketplace/structures"
)

type Auth interface {
	CreateUser(user structures.User) (int, error)
}

type Item interface {
}

type Service struct {
	Auth
	Item
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repos.Auth),
	}
}
