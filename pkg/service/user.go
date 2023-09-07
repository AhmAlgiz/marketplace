package service

import (
	"github.com/AhmAlgiz/marketplace/pkg/repository"
	"github.com/AhmAlgiz/marketplace/structures"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) UpdateUser(updateUser structures.UpdateUser, userId int) error {
	return s.repo.UpdateUser(updateUser, userId)
}

func (s *UserService) GetUserById(id int) ([]structures.GetUser, error) {
	return s.repo.GetUserById(id)
}
