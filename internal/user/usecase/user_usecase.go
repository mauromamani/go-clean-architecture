package usecase

import (
	"context"

	"github.com/mauromamani/go-clean-architecture/internal/user"
	"github.com/mauromamani/go-clean-architecture/internal/user/dtos"
	"github.com/mauromamani/go-clean-architecture/internal/user/entity"
)

type userUseCase struct {
	userRepository user.Repository
}

func NewUserUseCase(userRepo user.Repository) user.UseCase {
	return &userUseCase{userRepository: userRepo}
}

// GetUser:
func (uc *userUseCase) GetUsers(ctx context.Context) ([]*entity.User, error) {
	users, err := uc.userRepository.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetUserById:
func (uc *userUseCase) GetUserById(ctx context.Context, id int64) (*entity.User, error) {
	user, err := uc.userRepository.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByEmail:
func (uc *userUseCase) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	user, err := uc.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// CreateUser:
func (uc *userUseCase) CreateUser(ctx context.Context, user *dtos.CreateUserDto) (*entity.User, error) {
	u, err := uc.userRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// UpdateUser:
func (uc *userUseCase) UpdateUser(ctx context.Context, id int64, user *dtos.UpdateUserDto) (*entity.User, error) {
	updatedUser, err := uc.userRepository.UpdateUser(ctx, id, user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

// DeleteUser:
func (uc *userUseCase) DeleteUser(ctx context.Context, id int64) error {
	err := uc.userRepository.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
