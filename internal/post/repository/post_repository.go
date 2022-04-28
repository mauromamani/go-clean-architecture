package repository

import (
	"context"
	"database/sql"
	"time"

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
func (r *postRepository) GetPosts(ctx context.Context) ([]*entity.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	rows, err := r.DB.QueryContext(ctx, getPostsQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []*entity.Post
	for rows.Next() {
		var p entity.Post
		err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Body,
			&p.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		posts = append(posts, &p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

// GetPostById:
func (r *postRepository) GetPostById(ctx context.Context, id int64) (*entity.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var p entity.Post
	err := r.DB.QueryRowContext(ctx, getPostByIdQuery, id).Scan(
		&p.ID,
		&p.Title,
		&p.Body,
		&p.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

// CreatePost:
func (r *postRepository) CreatePost(ctx context.Context, user *entity.Post) (*entity.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

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
