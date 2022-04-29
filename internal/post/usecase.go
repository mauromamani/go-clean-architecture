package post

import (
	"context"

	"github.com/mauromamani/go-clean-architecture/internal/post/entity"
)

type UseCase interface {
	GetPosts(ctx context.Context) ([]*entity.Post, error)
	GetPostById(ctx context.Context, id int64) (*entity.Post, error)
	CreatePost(ctx context.Context, post *entity.Post) (*entity.Post, error)
	UpdatePost(ctx context.Context, id int64, post *entity.Post) (*entity.Post, error)
	DeletePost(ctx context.Context, id int64) error
}
