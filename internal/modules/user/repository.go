package user

import (
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindAll() ([]User, error) {
	rows, err := r.DB.Query(FindAllUsersQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Age, &user.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) Create(user User) (User, error) {
	err := r.DB.QueryRow(
		CreateUserQuery,
		user.FirstName, user.LastName, user.Email, user.Age, user.Password,
	).Scan(&user.ID)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
