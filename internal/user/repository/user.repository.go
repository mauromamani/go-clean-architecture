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
func (r *userRepository) Get(ctx context.Context) ([]*ent.User, error) {
	users, err := r.client.User.
		Query().
		All(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetById
func (r *userRepository) GetById(ctx context.Context, id int) (*ent.User, error) {
	u, err := r.client.User.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	return u, nil
}

// Create
func (r *userRepository) Create(ctx context.Context, user *ent.User) (*ent.User, error) {
	u, err := r.client.User.
		Create().
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
func (r *userRepository) Update(ctx context.Context, user *ent.User) (*ent.User, error) {
	userUpdated, err := r.client.User.
		UpdateOneID(user.ID).
		SetName(user.Name).
		SetEmail(user.Email).
		SetPassword(user.Password).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return userUpdated, nil
}

// Delete
func (r *userRepository) Delete(ctx context.Context, id int) error {
	err := r.client.User.
		DeleteOneID(id).
		Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
