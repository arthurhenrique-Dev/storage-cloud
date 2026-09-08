package services

import (
	"storage-api/modules/items/models"
	"storage-api/modules/items/repositories"
	"storage-api/shared/operations"
)

type ItemService struct {
	repo *repositories.ItemRepository
}

func NewItemService(repo *repositories.ItemRepository) *ItemService {
	return &ItemService{
		repo: repo,
	}
}

func (s *ItemService) Create(item *models.Item) error {
	return operations.NewPipeline("Create Item").
		Step("Save Item", func() error {
			return s.repo.Create(item)
		}).
		Execute()
}

func (s *ItemService) GetByID(id int64) (*models.Item, error) {
	var item *models.Item
	err := operations.NewPipeline("Get Item By ID").
		Step("Find Item", func() error {
			var err error
			item, err = s.repo.GetByID(id)
			return err
		}).
		Execute()

	return item, err
}

func (s *ItemService) List() ([]*models.Item, error) {
	var items []*models.Item
	err := operations.NewPipeline("List Items").
		Step("Fetch Items", func() error {
			var err error
			items, err = s.repo.List()
			return err
		}).
		Execute()

	return items, err
}

func (s *ItemService) Update(item *models.Item) error {
	return operations.NewPipeline("Update Item").
		Step("Save Item Update", func() error {
			return s.repo.Update(item)
		}).
		Execute()
}

func (s *ItemService) Delete(id int64) error {
	return operations.NewPipeline("Delete Item").
		Step("Save Item Delete", func() error {
			return s.repo.Delete(id)
		}).
		Execute()
}
