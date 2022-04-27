package repository

import (
	"context"
	"database/sql"

	"github.com/mauromamani/go-clean-architecture/internal/post"
	"github.com/mauromamani/go-clean-architecture/internal/post/entity"
)

type postRepository struct {
	DB *sql.DB
}

func NewPostRepository(DB *sql.DB) post.Repository {
	return &postRepository{DB: DB}
}

// GetPosts:
func (r *postRepository) GetPosts(ctx context.Context) ([]entity.Post, error) {
	return nil, nil
}

// GetPostById:
func (r *postRepository) GetPostById(ctx context.Context, id int64) (*entity.Post, error) {
	return nil, nil
}

// CreatePost:
func (r *postRepository) CreatePost(ctx context.Context, user *entity.Post) (*entity.Post, error) {
	return nil, nil
}

// UpdatePost:
func (r *postRepository) UpdatePost(ctx context.Context, id int64, user *entity.Post) (*entity.Post, error) {
	return nil, nil
}

// DeletePost:
func (r *postRepository) DeletePost(ctx context.Context, id int64) error {
	return nil
}
