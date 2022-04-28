package post

import (
	"context"

	"github.com/mauromamani/go-clean-architecture/internal/post/entity"
)

type Repository interface {
	GetPosts(ctx context.Context) ([]*entity.Post, error)
	GetPostById(ctx context.Context, id int64) (*entity.Post, error)
	CreatePost(ctx context.Context, user *entity.Post) (*entity.Post, error)
	UpdatePost(ctx context.Context, id int64, user *entity.Post) (*entity.Post, error)
	DeletePost(ctx context.Context, id int64) error
}
