package usecase

import (
	"context"
	"net/http"

	"github.com/mauromamani/go-clean-architecture/ent"
	"github.com/mauromamani/go-clean-architecture/internal/user"
	"github.com/mauromamani/go-clean-architecture/internal/user/dtos"
	httpErrors "github.com/mauromamani/go-clean-architecture/pkg/errors"
)

type userUseCase struct {
	userRepository user.Repository
}

func NewUserUseCase(userRepo user.Repository) user.UseCase {
	return &userUseCase{userRepository: userRepo}
}

// Get
func (u *userUseCase) Get(ctx context.Context) ([]*ent.User, error) {
	users, err := u.userRepository.Get(ctx)
	if err != nil {
		return nil, httpErrors.NewInternalServerError(err)
	}

	return users, nil
}

// GetById
func (u *userUseCase) GetById(ctx context.Context, id int) (*ent.User, error) {
	user, err := u.userRepository.GetById(ctx, id)
	if err != nil {
		return nil, httpErrors.NewInternalServerError(err)
	}

	return user, nil
}

// GetByEmail
func (u *userUseCase) GetByEmail(ctx context.Context, email string) (*ent.User, error) {
	user, err := u.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return nil, httpErrors.NewInternalServerError(err)
	}

	return user, nil
}

// Create
func (u *userUseCase) Create(ctx context.Context, user *dtos.CreateUserDto) (*ent.User, error) {
	existsUser, err := u.userRepository.GetByEmail(ctx, user.Email)
	if existsUser != nil || err == nil {
		return nil, httpErrors.NewRestError(http.StatusBadRequest, httpErrors.ErrEmailAlreadyExists, nil)
	}

	newUser, err := u.userRepository.Create(ctx, user)
	if err != nil {
		return nil, httpErrors.NewInternalServerError(err)
	}

	return newUser, nil
}

// Update
func (u *userUseCase) Update(ctx context.Context, user *ent.User) (*ent.User, error) {
	updatedUser, err := u.userRepository.Update(ctx, user)
	if err != nil {
		return nil, httpErrors.NewInternalServerError(err)
	}

	return updatedUser, nil
}

// Delete
func (u *userUseCase) Delete(ctx context.Context, id int) error {
	err := u.userRepository.Delete(ctx, id)
	if err != nil {
		return httpErrors.NewInternalServerError(err)
	}

	return nil
}
