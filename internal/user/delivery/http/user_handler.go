package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mauromamani/go-clean-architecture/internal/user"
	"github.com/mauromamani/go-clean-architecture/internal/user/dtos"
	httpErrors "github.com/mauromamani/go-clean-architecture/pkg/errors"
	"github.com/mauromamani/go-clean-architecture/pkg/utils"
)

type userHandlers struct {
	useCase user.UseCase
}

func NewUserHandlers(useCase user.UseCase) user.Handlers {
	return &userHandlers{
		useCase: useCase,
	}
}

// GetUser:
func (h *userHandlers) GetUser(c *gin.Context) {
	ctx := c.Request.Context()

	users, err := h.useCase.GetUser(ctx)

	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(200, users)
}

// GetUserById:
func (h *userHandlers) GetUserById(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, "Bad id")
		return
	}

	user, err := h.useCase.GetUserById(ctx, id)

	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(200, user)
}

// CreateUser:
func (h *userHandlers) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()

	user := &dtos.CreateUserDto{}
	if err := utils.ReadRequest(c, user); err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	newUser, err := h.useCase.CreateUser(ctx, user)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(200, newUser)
}

// UpdateUser:
func (h *userHandlers) UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, "Error Update")
		return
	}

	user := &dtos.UpdateUserDto{}
	if err := utils.ReadRequest(c, user); err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	updatedUser, err := h.useCase.UpdateUser(ctx, id, user)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(200, updatedUser)
}

// DeleteUser:
func (h *userHandlers) DeleteUser(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, "Bad id")
		return
	}

	err = h.useCase.DeleteUser(ctx, id)

	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(200, "Deleted")
}
