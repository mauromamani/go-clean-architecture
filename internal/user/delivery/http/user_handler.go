package http

import (
	"log"
	"net/http"

	"github.com/mauromamani/go-clean-architecture/internal/user"
	"github.com/mauromamani/go-clean-architecture/internal/user/dtos"
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
func (h *userHandlers) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := h.useCase.GetUsers(ctx)
	if err != nil {
		log.Println(err)
		log.Println("Error: GetUsers.user_handler")
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, map[string]interface{}{"users": users}, nil)
	if err != nil {
		log.Println(err)
		log.Println("Error: WriteJSON.user_handler")
	}
}

// GetUserById:
func (h *userHandlers) GetUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := utils.ReadIDParam(r)
	if err != nil {
		log.Println(err)
		log.Println("Error: utils.ReadIDParam.user_handler")
		return
	}

	u, err := h.useCase.GetUserById(ctx, id)
	if err != nil {
		log.Println(err)
		log.Println("Error: h.useCase.GetUserById.user_handler")
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, map[string]interface{}{"user": u}, nil)
	if err != nil {
		log.Println(err)
		log.Println("Error: utils.WriteJSON.user_handler")
		return
	}
}

// CreateUser:
func (h *userHandlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user := &dtos.CreateUserDto{}

	err := utils.ReadJSON(w, r, user)
	if err != nil {
		log.Println(err)
		log.Println("Error: utils.ReadJSON.user_handler")
		return
	}

	u, err := h.useCase.CreateUser(ctx, user)
	if err != nil {
		log.Println(err)
		log.Println("Error: h.useCase.CreateUser.user_handler")
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, map[string]interface{}{"user": u}, nil)
	if err != nil {
		log.Println(err)
		log.Println("Error: utils.WriteJSON.user_handler")
	}
}

// UpdateUser:
func (h *userHandlers) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// id, err := strconv.Atoi(c.Param("id"))

	// if err != nil {
	// 	c.JSON(400, "Error Update")
	// 	return
	// }

	user := &dtos.UpdateUserDto{}
	// if err := utils.ReadRequest(c, user); err != nil {
	// 	c.JSON(httpErrors.ErrorResponse(err))
	// 	return
	// }

	updatedUser, _ := h.useCase.UpdateUser(ctx, 1, user)

	print(updatedUser)

	w.Write([]byte("update"))
}

// DeleteUser:
func (h *userHandlers) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// id, err := strconv.Atoi(c.Param("id"))

	// if err != nil {
	// 	c.JSON(400, "Bad id")
	// 	return
	// }

	h.useCase.DeleteUser(ctx, 1)

	w.Write([]byte("delete"))
}
