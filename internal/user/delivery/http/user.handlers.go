package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mauromamani/go-clean-architecture/ent"
	"github.com/mauromamani/go-clean-architecture/internal/user"
	"github.com/mauromamani/go-clean-architecture/internal/user/dtos"
	httpErrors "github.com/mauromamani/go-clean-architecture/pkg/errors"
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

	users, err := h.useCase.Get(ctx)

	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(200, users)
}

// GetById
func (h *userHandlers) GetById(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, "Bad id")
		return
	}

	user, err := h.useCase.GetById(ctx, id)

	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(200, user)
}

// Create
func (h *userHandlers) Create(c *gin.Context) {
	ctx := c.Request.Context()

	user := &dtos.CreateUserDto{}
	if err := c.Bind(user); err != nil {
		c.JSON(404, err)
	}

	newUser, err := h.useCase.Create(ctx, user)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(200, newUser)
}

// Update
func (h *userHandlers) Update(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, "Error Update")
		return
	}

	user := &ent.User{}
	if err := c.Bind(user); err != nil {
		c.JSON(404, err)
		return
	}
	user.ID = id

	updatedUser, err := h.useCase.Update(ctx, user)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(200, updatedUser)
}

// Delete
func (h *userHandlers) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, "Bad id")
		return
	}

	err = h.useCase.Delete(ctx, id)

	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(200, "Deleted")
}
