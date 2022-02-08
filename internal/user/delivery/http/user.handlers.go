package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mauromamani/go-clean-architecture/internal/user"
)

type userHandlers struct{}

func NewUserHandlers() user.Handlers {
	return &userHandlers{}
}

//
func (h *userHandlers) Get(c *gin.Context) {
	c.JSON(200, "Get")
}

//
func (h *userHandlers) GetById(c *gin.Context) {
	c.JSON(200, "GetById")

}

//
func (h *userHandlers) Create(c *gin.Context) {
	c.JSON(200, "Create")

}

//
func (h *userHandlers) Update(c *gin.Context) {
	c.JSON(200, "Update")
}

//
func (h *userHandlers) Delete(c *gin.Context) {
	c.JSON(200, "Delete")
}
