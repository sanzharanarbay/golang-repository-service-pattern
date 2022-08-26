package services

import (
	"github.com/sanzharanarbay/repository-service-pattern/models"
	"github.com/sanzharanarbay/repository-service-pattern/repositories"
)

type ItemService struct {
	itemRepository *repositories.ItemRepository
}

func NewItemService(itemRepository *repositories.ItemRepository) *ItemService {
	return &ItemService{
		itemRepository: itemRepository,
	}
}

func (s *ItemService) IsItemAvailable(id int) (bool, error) {
	item, err := s.itemRepository.GetItemById(id)
	if err == nil && item != nil {
		return true, nil
	}
	return false, err
}

func (s *ItemService) GetSingleItem(ID int) (*models.Item, error) {
	item, err := s.itemRepository.GetItemById(ID)
	return item, err
}

func (s *ItemService) GetAllItems() (*[]models.Item, error) {
	itemList, err := s.itemRepository.GetAllItems()
	return itemList, err
}

func (s *ItemService) InsertItem(item *models.Item) (bool, error) {
	state, err := s.itemRepository.SaveItem(item)
	return state, err
}

func (s *ItemService) DeleteItem(id int) (bool, error) {
	var err error
	found, err := s.IsItemAvailable(id)
	if found == false {
		return false, err
	}
	state, err := s.itemRepository.DeleteItem(id)
	return state, err
}

func (s *ItemService) UpdateItem(item *models.Item, ID int) (bool, error) {
	found, err := s.IsItemAvailable(ID)
	if found == false {
		return false, err
	}
	state, err := s.itemRepository.UpdateItem(item, ID)
	return state, err
}
