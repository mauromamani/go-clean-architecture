package http

import (
	"fmt"
	"net/http"

	"github.com/mauromamani/go-clean-architecture/internal/user"
	"github.com/mauromamani/go-clean-architecture/internal/user/dtos"
	httpErrors "github.com/mauromamani/go-clean-architecture/pkg/errors"
	"github.com/mauromamani/go-clean-architecture/pkg/logger"
	"github.com/mauromamani/go-clean-architecture/pkg/utils"
)

type userHandlers struct {
	useCase user.UseCase
	logger  logger.Logger
}

func NewUserHandlers(useCase user.UseCase, logger logger.Logger) user.Handlers {
	return &userHandlers{
		useCase: useCase,
		logger:  logger,
	}
}

// GetUser:
func (h *userHandlers) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := h.useCase.GetUsers(ctx)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, map[string]interface{}{"users": users}, nil)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
	}
}

// GetUserById:
func (h *userHandlers) GetUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := utils.ReadIDParam(r)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
		return
	}

	u, err := h.useCase.GetUserById(ctx, id)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, map[string]interface{}{"user": u}, nil)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
	}
}

// CreateUser:
func (h *userHandlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user := &dtos.CreateUserDto{}

	err := utils.ReadJSON(w, r, user)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
		return
	}

	u, err := h.useCase.CreateUser(ctx, user)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/api/v1/users/%d", u.ID))

	err = utils.WriteJSON(w, http.StatusOK, map[string]interface{}{"user": u}, headers)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
	}
}

// UpdateUser:
func (h *userHandlers) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := utils.ReadIDParam(r)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
		return
	}

	user := &dtos.UpdateUserDto{}

	err = utils.ReadJSON(w, r, user)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
		return
	}

	updatedUser, err := h.useCase.UpdateUser(ctx, id, user)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, map[string]interface{}{"user": updatedUser}, nil)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
	}
}

// DeleteUser:
func (h *userHandlers) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := utils.ReadIDParam(r)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
		return
	}

	err = h.useCase.DeleteUser(ctx, id)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, map[string]interface{}{"user": "user deleted!"}, nil)
	if err != nil {
		h.logger.Error(err.Error())
		status, restErr := httpErrors.ErrorResponse(err)
		utils.WriteJSON(w, status, map[string]interface{}{"error": restErr}, nil)
	}
}
