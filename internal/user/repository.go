package user

import "context"

type Repository interface {
	Get(ctx context.Context)
	GetById(ctx context.Context)
	Create(ctx context.Context)
	Update(ctx context.Context)
	Delete(ctx context.Context)
}
