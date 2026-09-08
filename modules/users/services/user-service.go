package services

import (
	"storage-api/modules/users/models"
	"storage-api/modules/users/repositories"
	"storage-api/shared/operations"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(user *models.User) error {
	return operations.NewPipeline("Create User").
		Step("Save User", func() error {
			return s.repo.Create(user)
		}).
		Execute()
}

func (s *UserService) GetByID(id int64) (*models.User, error) {
	var user *models.User
	err := operations.NewPipeline("Get User By ID").
		Step("Find User", func() error {
			var err error
			user, err = s.repo.GetByID(id)
			return err
		}).
		Execute()

	return user, err
}

func (s *UserService) GetByEmail(email string) (*models.User, error) {
	var user *models.User
	err := operations.NewPipeline("Get User By Email").
		Step("Find User By Email", func() error {
			var err error
			user, err = s.repo.GetByEmail(email)
			return err
		}).
		Execute()

	return user, err
}

func (s *UserService) List() ([]*models.User, error) {
	var users []*models.User
	err := operations.NewPipeline("List Users").
		Step("Fetch Users", func() error {
			var err error
			users, err = s.repo.List()
			return err
		}).
		Execute()

	return users, err
}

func (s *UserService) Delete(id int64) error {
	return operations.NewPipeline("Delete User").
		Step("Save User Delete", func() error {
			return s.repo.Delete(id)
		}).
		Execute()
}
