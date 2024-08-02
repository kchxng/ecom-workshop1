package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service UserServiceInterface
}

func NewUserHandler(router *gin.RouterGroup, service UserServiceInterface) UserHandlerInterface {
	handler := &userHandler{
		service: service,
	}

	// Route path
	router.GET("/", handler.GetAll)
	router.GET("/trash", handler.GetAllTrash)
	return handler
}

// GetAll implements CustomHandlerInterface.
func (h *userHandler) GetAll(c *gin.Context) {
	data, err := h.service.GetAll(0, 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   data,
	})
}

// GetAllTrash implements CustomHandlerInterface.
func (h *userHandler) GetAllTrash(c *gin.Context) {
	panic("unimplemented")
}

// GetSingle implements CustomHandlerInterface.
func (h *userHandler) GetSingle(c *gin.Context) {
	panic("unimplemented")
}

// Create implements CustomHandlerInterface.
func (h *userHandler) Create(c *gin.Context) {
	panic("unimplemented")
}

// Delete implements CustomHandlerInterface.
func (h *userHandler) Delete(c *gin.Context) {
	panic("unimplemented")
}

// Destroy implements CustomHandlerInterface.
func (h *userHandler) Destroy(c *gin.Context) {
	panic("unimplemented")
}

// Restore implements CustomHandlerInterface.
func (*userHandler) Restore(*gin.Context) {
	panic("unimplemented")
}

// Update implements CustomHandlerInterface.
func (*userHandler) Update(*gin.Context) {
	panic("unimplemented")
}
