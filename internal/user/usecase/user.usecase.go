package usecase

import (
	"context"

	"github.com/mauromamani/go-clean-architecture/ent"
	"github.com/mauromamani/go-clean-architecture/internal/user"
)

type userUseCase struct {
	userRepository user.Repository
}

func NewUserUseCase(userRepo user.Repository) user.UseCase {
	return &userUseCase{userRepository: userRepo}
}

// Get
func (u *userUseCase) Get(ctx context.Context) {
	u.userRepository.Get(ctx)
}

// GetById
func (u *userUseCase) GetById(ctx context.Context) {
	u.userRepository.GetById(ctx)
}

// Create
func (u *userUseCase) Create(ctx context.Context, user *ent.User) (*ent.User, error) {
	newUser, err := u.userRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

// Update
func (u *userUseCase) Update(ctx context.Context) {
	u.userRepository.Update(ctx)
}

// Delete
func (u *userUseCase) Delete(ctx context.Context) {
	u.userRepository.Delete(ctx)
}
