package app

import (
	"github.com/kchxng/ecom-api/models"

	"gorm.io/gorm"
)

type customRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &customRepository{db: db}
}

// Create implements CustomRepositoryInterface.
func (*customRepository) Create(body *models.Users) (*string, error) {
	panic("unimplemented")
}

// Delete implements CustomRepositoryInterface.
func (*customRepository) Delete(id string) error {
	panic("unimplemented")
}

// Destroy implements CustomRepositoryInterface.
func (*customRepository) Destroy(id string) error {
	panic("unimplemented")
}

// GetAll implements CustomRepositoryInterface.
func (*customRepository) GetAll(page int, pageSize int) ([]models.Users, error) {
	panic("unimplemented")
}

// GetAllTrash implements CustomRepositoryInterface.
func (*customRepository) GetAllTrash(page int, pageSize int) ([]models.Users, error) {
	panic("unimplemented")
}

// GetSingle implements CustomRepositoryInterface.
func (*customRepository) GetSingle(id string) (*models.Users, error) {
	panic("unimplemented")
}

// Restore implements CustomRepositoryInterface.
func (*customRepository) Restore(id string) error {
	panic("unimplemented")
}

// Update implements CustomRepositoryInterface.
func (*customRepository) Update(id string, body *models.Users) error {
	panic("unimplemented")
}
