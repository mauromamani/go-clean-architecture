package repository

const (
	getUsersQuery = `
		SELECT id, name, email, created_at
		FROM users
	`

	getUserByIdQuery = `
		SELECT id, name, email, created_at
		FROM users
		WHERE id = $1
	`

	createUserQuery = `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING *
	`

	updateUserQuery = `
		UPDATE users
		SET 
			name = COALESCE(NULLIF($1, ''), name),
			email = COALESCE(NULLIF($2, ''), email),
			password = COALESCE(NULLIF($3, ''), password)
		WHERE id = $4
		RETURNING *
	`

	deleteUserQuery = `
		DELETE FROM users
		WHERE id = $1
	`
)
