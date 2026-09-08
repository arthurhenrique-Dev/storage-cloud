package services

import (
	"storage-api/modules/suppliers/models"
	"storage-api/modules/suppliers/repositories"
	"storage-api/shared/operations"
)

type SupplierService struct {
	repo *repositories.SupplierRepository
}

func NewSupplierService(repo *repositories.SupplierRepository) *SupplierService {
	return &SupplierService{
		repo: repo,
	}
}

func (s *SupplierService) Create(supplier *models.Supplier) error {
	return operations.NewPipeline("Create Supplier").
		Step("Save Supplier", func() error {
			return s.repo.Create(supplier)
		}).
		Execute()
}

func (s *SupplierService) GetByID(id int64) (*models.Supplier, error) {
	var supplier *models.Supplier
	err := operations.NewPipeline("Get Supplier By ID").
		Step("Find Supplier", func() error {
			var err error
			supplier, err = s.repo.GetByID(id)
			return err
		}).
		Execute()

	return supplier, err
}

func (s *SupplierService) List() ([]*models.Supplier, error) {
	var suppliers []*models.Supplier
	err := operations.NewPipeline("List Suppliers").
		Step("Fetch Suppliers", func() error {
			var err error
			suppliers, err = s.repo.List()
			return err
		}).
		Execute()

	return suppliers, err
}

func (s *SupplierService) Update(supplier *models.Supplier) error {
	return operations.NewPipeline("Update Supplier").
		Step("Save Supplier Update", func() error {
			return s.repo.Update(supplier)
		}).
		Execute()
}

func (s *SupplierService) Delete(id int64) error {
	return operations.NewPipeline("Delete Supplier").
		Step("Save Supplier Delete", func() error {
			return s.repo.Delete(id)
		}).
		Execute()
}
