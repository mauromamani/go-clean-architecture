package repository

import "github.com/mauromamani/go-clean-architecture/internal/user"

type userRepository struct{}

func NewUserRepository() user.Repository {
	return &userRepository{}
}

// Get
func (r *userRepository) Get() {

}

// GetById
func (r *userRepository) GetById() {

}

// Create
func (r *userRepository) Create() {

}

// Update
func (r *userRepository) Update() {

}

// Delete
func (r *userRepository) Delete() {

}
