package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/mauromamani/go-clean-architecture/internal/post"
	"github.com/mauromamani/go-clean-architecture/internal/post/dtos"
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
			&p.UserID,
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
		&p.UserID,
	)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

// CreatePost:
func (r *postRepository) CreatePost(ctx context.Context, post *dtos.CreatePostDto) (*entity.Post, error) {
	args := []interface{}{post.Title, post.Body, post.UserID}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var p entity.Post

	err := r.DB.QueryRowContext(ctx, createPostQuery, args...).Scan(
		&p.ID,
		&p.CreatedAt,
		&p.Title,
		&p.Body,
		&p.UserID,
	)

	// TODO: Error cuando se usa un user_id que  no existe
	if err != nil {
		return nil, err
	}

	return &p, nil
}

// UpdatePost:
func (r *postRepository) UpdatePost(ctx context.Context, id int64, post *dtos.UpdatePostDto) (*entity.Post, error) {
	args := []interface{}{post.Title, post.Title, id}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var p entity.Post

	err := r.DB.QueryRowContext(ctx, updatePostQuery, args...).Scan(
		&p.ID,
		&p.CreatedAt,
		&p.Title,
		&p.Body,
		&p.UserID,
	)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

// DeletePost:
func (r *postRepository) DeletePost(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := r.DB.ExecContext(ctx, deletePostQuery, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// TODO: Agregar el error ErrRecordNotFound
	if rowsAffected == 0 {
		return errors.New("record not found")
	}

	return nil
}
