package repository

import (
	"context"

	"github.com/mauromamani/go-clean-architecture/ent"
	"github.com/mauromamani/go-clean-architecture/internal/user"
)

type userRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) user.Repository {
	return &userRepository{
		client: client,
	}
}

// Get
func (r *userRepository) Get(ctx context.Context) {

}

// GetById
func (r *userRepository) GetById(ctx context.Context) {

}

// Create
func (r *userRepository) Create(ctx context.Context, user *ent.User) (*ent.User, error) {
	u, err := r.client.User.Create().
		SetName(user.Name).
		SetPassword(user.Password).
		SetEmail(user.Email).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return u, nil
}

// Update
func (r *userRepository) Update(ctx context.Context) {

}

// Delete
func (r *userRepository) Delete(ctx context.Context) {

}
