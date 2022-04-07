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

	deleteUserByIdQuery = `
		DELETE FROM users
		WHERE id = $1
	`
)
