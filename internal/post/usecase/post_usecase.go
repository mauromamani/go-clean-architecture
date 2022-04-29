package usecase

import (
	"context"

	"github.com/mauromamani/go-clean-architecture/internal/post"
	"github.com/mauromamani/go-clean-architecture/internal/post/entity"
)

type postUseCase struct {
	postRepository post.Repository
}

func NewPostUseCase(postRepo post.Repository) post.UseCase {
	return &postUseCase{
		postRepository: postRepo,
	}
}

// GetPosts:
func (uc *postUseCase) GetPosts(ctx context.Context) ([]*entity.Post, error) {
	return nil, nil
}

// GetPostById:
func (uc *postUseCase) GetPostById(ctx context.Context, id int64) (*entity.Post, error) {
	return nil, nil
}

// CreatePost:
func (uc *postUseCase) CreatePost(ctx context.Context, post *entity.Post) (*entity.Post, error) {
	return nil, nil
}

// UpdatePost:
func (uc *postUseCase) UpdatePost(ctx context.Context, id int64, post *entity.Post) (*entity.Post, error) {
	return nil, nil
}

// DeletePost:
func (uc *postUseCase) DeletePost(ctx context.Context, id int64) error {
	return nil
}
