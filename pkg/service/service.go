package service

import "github.com/AhmAlgiz/marketplace/pkg/repository"

type Auth interface {
}

type Item interface {
}

type Service struct {
	Auth
	Item
}

func NewService(rep *repository.Repository) *Service {
	return &Service{}
}
