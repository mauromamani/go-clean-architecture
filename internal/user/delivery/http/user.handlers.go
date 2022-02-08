package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mauromamani/go-clean-architecture/internal/user"
)

type userHandlers struct {
	useCase user.UseCase
}

func NewUserHandlers(userUC user.UseCase) user.Handlers {
	return &userHandlers{useCase: userUC}
}

// Get
func (h *userHandlers) Get(c *gin.Context) {
	c.JSON(200, "Get")
}

// GetById
func (h *userHandlers) GetById(c *gin.Context) {
	c.JSON(200, "GetById")

}

// Create
func (h *userHandlers) Create(c *gin.Context) {
	c.JSON(200, "Create")

}

// Update
func (h *userHandlers) Update(c *gin.Context) {
	c.JSON(200, "Update")
}

// Delete
func (h *userHandlers) Delete(c *gin.Context) {
	c.JSON(200, "Delete")
}
