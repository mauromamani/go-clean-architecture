package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mauromamani/go-clean-architecture/internal/user"
	"github.com/mauromamani/go-clean-architecture/internal/user/dtos"
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
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, map[string]interface{}{"users": users}, nil)
	if err != nil {
		h.logger.Error(err.Error())
	}
}

// GetUserById:
func (h *userHandlers) GetUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := utils.ReadIDParam(r)
	if err != nil {
		h.logger.Error(err.Error())
		return
	}

	u, err := h.useCase.GetUserById(ctx, id)
	if err != nil {
		h.logger.Error(err.Error())
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, map[string]interface{}{"user": u}, nil)
	if err != nil {
		h.logger.Error(err.Error())
		return
	}
}

// CreateUser:
func (h *userHandlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user := &dtos.CreateUserDto{}

	err := utils.ReadJSON(w, r, user)
	if err != nil {
		h.logger.Error(err.Error())
		return
	}

	u, err := h.useCase.CreateUser(ctx, user)
	if err != nil {
		h.logger.Error(err.Error())
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/api/v1/users/%d", u.ID))

	err = utils.WriteJSON(w, http.StatusOK, map[string]interface{}{"user": u}, headers)
	if err != nil {
		h.logger.Error(err.Error())
	}
}

// UpdateUser:
func (h *userHandlers) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := utils.ReadIDParam(r)
	if err != nil {
		h.logger.Error(err.Error())
		log.Println(err)
		log.Println("Error: utils.ReadIDParam.user_handler")
		return
	}

	user := &dtos.UpdateUserDto{}

	err = utils.ReadJSON(w, r, user)
	if err != nil {
		h.logger.Error(err.Error())
		return
	}

	updatedUser, err := h.useCase.UpdateUser(ctx, id, user)
	if err != nil {
		h.logger.Error(err.Error())
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, map[string]interface{}{"user": updatedUser}, nil)
	if err != nil {
		h.logger.Error(err.Error())
	}
}

// DeleteUser:
func (h *userHandlers) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := utils.ReadIDParam(r)
	if err != nil {
		h.logger.Error(err.Error())
		return
	}

	err = h.useCase.DeleteUser(ctx, id)
	if err != nil {
		h.logger.Error(err.Error())
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, map[string]interface{}{"user": "user deleted!"}, nil)
	if err != nil {
		h.logger.Error(err.Error())
	}
}
