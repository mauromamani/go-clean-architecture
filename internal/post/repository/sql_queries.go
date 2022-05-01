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
		VALUES (TRIM($1), TRIM($2), $3)
		RETURNING *
	`

	updatePostQuery = `
		UPDATE posts
		SET
			title = TRIM(COALESCE(NULLIF($1, ''), title)),
			body = TRIM(COALESCE(NULLIF($2, ''), body))
		WHERE id = $3
		RETURNING *
			
	`

	deletePostQuery = `
		DELETE FROM posts
		WHERE id = $1
	`
)
