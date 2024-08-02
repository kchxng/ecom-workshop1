package app

import "github.com/gin-gonic/gin"

type UserHandlerInterface interface {
	GetAll(*gin.Context)
	GetAllTrash(*gin.Context)
	GetSingle(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	Destroy(*gin.Context)
	Restore(*gin.Context)
}

type UserServiceInterface interface {
	GetAll(page int, pageSize int) ([]Model, error)
	GetAllTrash(page int, pageSize int) ([]Model, error)
	GetSingle(id string) (*Model, error)
	Create(body *Model) (*string, error)
	Update(id string, body *Model) error
	Delete(id string) error
	Destroy(id string) error
	Restore(id string) error
}

type UserRepositoryInterface interface {
	GetAll(page int, pageSize int) ([]Model, error)
	GetAllTrash(page int, pageSize int) ([]Model, error)
	GetSingle(id string) (*Model, error)
	Create(body *Model) (*string, error)
	Update(id string, body *Model) error
	Delete(id string) error
	Destroy(id string) error
	Restore(id string) error
}
