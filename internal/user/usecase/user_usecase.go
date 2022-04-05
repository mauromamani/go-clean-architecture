package usecase

import (
	"context"
	"net/http"

	"github.com/mauromamani/go-clean-architecture/internal/user"
	"github.com/mauromamani/go-clean-architecture/internal/user/dtos"
	"github.com/mauromamani/go-clean-architecture/internal/user/entity"
	httpErrors "github.com/mauromamani/go-clean-architecture/pkg/errors"
)

type userUseCase struct {
	userRepository user.Repository
}

func NewUserUseCase(userRepo user.Repository) user.UseCase {
	return &userUseCase{userRepository: userRepo}
}

// GetUser
func (u *userUseCase) GetUser(ctx context.Context) (*entity.User, error) {
	users, err := u.userRepository.GetUser(ctx)
	if err != nil {
		return nil, httpErrors.NewInternalServerError(err)
	}

	return users, nil
}

// GetUserById
func (u *userUseCase) GetUserById(ctx context.Context, id int) (*entity.User, error) {
	user, err := u.userRepository.GetUserById(ctx, id)
	if err != nil {
		return nil, httpErrors.NewInternalServerError(err)
	}

	return user, nil
}

// GetUserByEmail
func (u *userUseCase) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	user, err := u.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, httpErrors.NewInternalServerError(err)
	}

	return user, nil
}

// CreateUser
func (u *userUseCase) CreateUser(ctx context.Context, user *dtos.CreateUserDto) (*entity.User, error) {
	existsUser, err := u.userRepository.GetUserByEmail(ctx, user.Email)
	if existsUser != nil || err == nil {
		return nil, httpErrors.NewRestError(http.StatusBadRequest, httpErrors.ErrEmailAlreadyExists, nil)
	}

	newUser, err := u.userRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, httpErrors.NewInternalServerError(err)
	}

	return newUser, nil
}

// UpdateUser
func (u *userUseCase) UpdateUser(ctx context.Context, id int, user *dtos.UpdateUserDto) (*entity.User, error) {
	updatedUser, err := u.userRepository.UpdateUser(ctx, id, user)
	if err != nil {
		return nil, httpErrors.NewInternalServerError(err)
	}

	return updatedUser, nil
}

// DeleteUser
func (u *userUseCase) DeleteUser(ctx context.Context, id int) error {
	err := u.userRepository.DeleteUser(ctx, id)
	if err != nil {
		return httpErrors.NewInternalServerError(err)
	}

	return nil
}
