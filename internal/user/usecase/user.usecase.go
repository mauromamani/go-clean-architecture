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
func (u *userUseCase) Get(ctx context.Context) ([]*ent.User, error) {
	return u.userRepository.Get(ctx)
}

// GetById
func (u *userUseCase) GetById(ctx context.Context, id int) (*ent.User, error) {
	return u.userRepository.GetById(ctx, id)
}

// Create
func (u *userUseCase) Create(ctx context.Context, user *ent.User) (*ent.User, error) {
	return u.userRepository.Create(ctx, user)
}

// Update
func (u *userUseCase) Update(ctx context.Context, user *ent.User) (*ent.User, error) {
	return u.userRepository.Update(ctx, user)
}

// Delete
func (u *userUseCase) Delete(ctx context.Context, id int) error {
	return u.userRepository.Delete(ctx, id)
}
