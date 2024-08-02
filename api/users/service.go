package app

import (
	"github.com/kchxng/ecom-api/models"
	"time"
)

type userService struct {
	repo UserRepositoryInterface
}

func NewUserService(repo UserRepositoryInterface) UserServiceInterface {
	return &userService{repo: repo}
}

// GetAll implements CustomServiceInterface.
func (s *userService) GetAll(page int, pageSize int) ([]Model, error) {
	now := time.Now()
	email := "chengjs2018@gmail.com"
	data := Model{
		Id:        1,
		Username:  "cheng",
		Password:  "<PASSWORD>",
		Email:     &email,
		CreatedAt: &now,
		UpdatedAt: &now,
		DeletedAt: nil,
	}

	return []Model{data}, nil
}

// GetAllTrash implements CustomServiceInterface.
func (s *userService) GetAllTrash(page int, pageSize int) ([]models.Users, error) {
	panic("unimplemented")
}

// GetSingle implements CustomServiceInterface.
func (s *userService) GetSingle(id string) (*models.Users, error) {
	panic("unimplemented")
}

// Create implements CustomServiceInterface.
func (s *userService) Create(body *models.Users) (*string, error) {
	panic("unimplemented")
}

// Delete implements CustomServiceInterface.
func (s *userService) Delete(id string) error {
	panic("unimplemented")
}

// Destroy implements CustomServiceInterface.
func (s *userService) Destroy(id string) error {
	panic("unimplemented")
}

// Restore implements CustomServiceInterface.
func (s *userService) Restore(id string) error {
	panic("unimplemented")
}

// Update implements CustomServiceInterface.
func (s *userService) Update(id string, body *models.Users) error {
	panic("unimplemented")
}
