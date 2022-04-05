package http

import (
	"net/http"

	"github.com/mauromamani/go-clean-architecture/internal/user"
	"github.com/mauromamani/go-clean-architecture/internal/user/dtos"
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
func (h *userHandlers) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := h.useCase.GetUser(ctx)

	if err != nil {
		return
	}

	print(users)

	w.Write([]byte("hello"))
}

// GetUserById:
func (h *userHandlers) GetUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// id, err := strconv.Atoi(c.Param("id"))

	// if err != nil {
	// 	return
	// }

	user, err := h.useCase.GetUserById(ctx, 2)
	print(user)

	if err != nil {
		return
	}

	w.Write([]byte("getuserbiyd"))
}

// CreateUser:
func (h *userHandlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user := &dtos.CreateUserDto{}
	// if err := utils.ReadRequest(c, user); err != nil {
	// 	c.JSON(httpErrors.ErrorResponse(err))
	// 	return
	// }

	newUser, _ := h.useCase.CreateUser(ctx, user)
	print(newUser)

	w.Write([]byte("CREATE user"))
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
