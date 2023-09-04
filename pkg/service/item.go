package service

import (
	"github.com/AhmAlgiz/marketplace/pkg/repository"
	"github.com/AhmAlgiz/marketplace/structures"
)

type ItemService struct {
	repo repository.Item
}

func NewItemService(repo repository.Item) *ItemService {
	return &ItemService{repo: repo}
}

func (s *ItemService) CreateItem(input structures.Item) (int, error) {
	return s.repo.CreateItem(input)
}

func (s *ItemService) GetItemById(id int) ([]structures.Item, error) {
	return s.repo.GetItemById(id)
}

func (s *ItemService) GetItemByTitle(title string) ([]structures.Item, error) {
	return s.repo.GetItemByTitle(title)
}

func (s *ItemService) GetItemByUsername(username string) ([]structures.Item, error) {
	return s.repo.GetItemByUsername(username)
}

func (s *ItemService) DeleteItem(id, userId int) error {
	return s.repo.DeleteItem(id, userId)
}

func (s *ItemService) GetAllItems() ([]structures.Item, error) {
	return s.repo.GetAllItems()
}
