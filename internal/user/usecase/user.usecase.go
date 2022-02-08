package usecase

import "github.com/mauromamani/go-clean-architecture/internal/user"

type userUseCase struct{}

func NewUserUseCase() user.UseCase {
	return &userUseCase{}
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
