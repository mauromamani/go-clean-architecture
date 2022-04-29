package repository

const (
	getPostsQuery = `
		SELECT id, title, body, created_at, user_id
		FROM posts
	`

	getPostByIdQuery = `
		SELECT id, title, body, created_at, user_id
		FROM posts
		WHERE id = $1
	`

	createPostQuery = `
		INSERT INTO posts (title, body, user_id)
		VALUES ($1, $2, $3)
		RETURNING *
	`

	updatePostQuery = `
		UPDATE posts
		SET
			title = COALESCE(NULLIF($1, ''), title),
			body = COALESCE(NULLIF($2, ''), body)
		WHERE id = $3
		RETURNING *
			
	`

	deletePostQuery = `
		DELETE FROM posts
		WHERE id = $1
	`
)
