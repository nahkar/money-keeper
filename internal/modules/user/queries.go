package user

const (
	FindAllUsersQuery = `
		SELECT user_id, first_name, last_name, email, age
		FROM users
	`

	FindUserByIDQuery = `
		SELECT user_id, first_name, last_name, email, age, created_at
		FROM users 
		WHERE user_id = $1
	`

	CreateUserQuery = `
		INSERT INTO users (first_name, last_name, email, age, password) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING user_id
	`
)
