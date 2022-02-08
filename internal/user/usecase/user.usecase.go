package usecase

import "github.com/mauromamani/go-clean-architecture/internal/user"

type userUseCase struct {
	userRepository user.Repository
}

func NewUserUseCase(userRepo user.Repository) user.UseCase {
	return &userUseCase{userRepository: userRepo}
}

// Get
func (u *userUseCase) Get() {

}

// GetById
func (u *userUseCase) GetById() {

}

// Create
func (u *userUseCase) Create() {

}

// Update
func (u *userUseCase) Update() {

}

// Delete
func (u *userUseCase) Delete() {

}
