package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mauromamani/go-clean-architecture/ent"
	"github.com/mauromamani/go-clean-architecture/internal/user"
)

type userHandlers struct {
	useCase user.UseCase
}

func NewUserHandlers(useCase user.UseCase) user.Handlers {
	return &userHandlers{
		useCase: useCase,
	}
}

// Get
func (h *userHandlers) Get(c *gin.Context) {
	ctx := c.Request.Context()

	h.useCase.Get(ctx)

	c.JSON(200, "Get")
}

// GetById
func (h *userHandlers) GetById(c *gin.Context) {
	ctx := c.Request.Context()

	h.useCase.GetById(ctx)

	c.JSON(200, "GetById")
}

// Create
func (h *userHandlers) Create(c *gin.Context) {
	ctx := c.Request.Context()

	user := &ent.User{}
	if err := c.Bind(user); err != nil {
		c.JSON(404, err)
	}

	newUser, err := h.useCase.Create(ctx, user)
	if err != nil {
		c.JSON(404, err)
		return
	}

	c.JSON(200, newUser)
}

// Update
func (h *userHandlers) Update(c *gin.Context) {
	ctx := c.Request.Context()

	h.useCase.Update(ctx)

	c.JSON(200, "Update")
}

// Delete
func (h *userHandlers) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	h.useCase.Delete(ctx)

	c.JSON(200, "Delete")
}
