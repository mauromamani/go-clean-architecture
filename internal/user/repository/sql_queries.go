package repository

const (
	getUsersQuery = `
		SELECT id, name, email, created_at
		FROM users
	`
	createUserQuery = `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING *
	`
)
