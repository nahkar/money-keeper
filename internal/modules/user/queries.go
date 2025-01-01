package user

const (
	FindAllUsersQuery = `
		SELECT user_id, first_name, last_name, email, age
		FROM users
	`

	FindUserByIDQuery = `
		SELECT user_id, first_name, last_name, email, age
		FROM users 
		WHERE user_id = $1
	`

	CreateUserQuery = `
		INSERT INTO users (first_name, last_name, email, age, password) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING user_id
	`

	UpdateUserQuery = `
		UPDATE users 
		SET first_name = $1, last_name = $2, email = $3, age = $4, password = $5
		WHERE user_id = $6
	`

	DeleteUserQuery = `
		DELETE FROM users 
		WHERE user_id = $1
	`
)
