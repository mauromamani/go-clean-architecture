package usecase

import (
	"context"

	"github.com/mauromamani/go-clean-architecture/internal/post"
	"github.com/mauromamani/go-clean-architecture/internal/post/dtos"
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
	posts, err := uc.postRepository.GetPosts(ctx)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// GetPostById:
func (uc *postUseCase) GetPostById(ctx context.Context, id int64) (*entity.Post, error) {
	post, err := uc.postRepository.GetPostById(ctx, id)
	if err != nil {
		return nil, err
	}

	return post, nil
}

// CreatePost:
func (uc *postUseCase) CreatePost(ctx context.Context, post *dtos.CreatePostDto) (*entity.Post, error) {
	p, err := uc.postRepository.CreatePost(ctx, post)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// UpdatePost:
func (uc *postUseCase) UpdatePost(ctx context.Context, id int64, post *dtos.UpdatePostDto) (*entity.Post, error) {
	p, err := uc.postRepository.UpdatePost(ctx, id, post)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// DeletePost:
func (uc *postUseCase) DeletePost(ctx context.Context, id int64) error {
	err := uc.postRepository.DeletePost(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
