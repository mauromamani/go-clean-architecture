package usecase

import (
	"context"

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
func (u *userUseCase) Create(ctx context.Context) {
	u.userRepository.Create(ctx)
}

// Update
func (u *userUseCase) Update(ctx context.Context) {
	u.userRepository.Update(ctx)
}

// Delete
func (u *userUseCase) Delete(ctx context.Context) {
	u.userRepository.Delete(ctx)
}
