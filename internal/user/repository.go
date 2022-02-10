package user

import (
	"context"

	"github.com/mauromamani/go-clean-architecture/ent"
)

type Repository interface {
	Get(ctx context.Context)
	GetById(ctx context.Context)
	Create(ctx context.Context, user *ent.User) (*ent.User, error)
	Update(ctx context.Context)
	Delete(ctx context.Context)
}
