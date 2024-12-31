package user

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (r *UserRepository) FindAll() ([]UserResponse, error) {
	rows, err := r.DB.Query(FindAllUsersQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []UserResponse

	for rows.Next() {
		var user UserResponse

		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Age); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) FindById(id int) (UserResponse, error) {
	var user UserResponse

	err := r.DB.QueryRow(FindUserByIDQuery, id).
		Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Age)

	if err != nil {
		if err == sql.ErrNoRows {
			return UserResponse{}, fmt.Errorf("%w: user with id %d", ErrUserNotFound, id)
		}
		return UserResponse{}, err
	}

	return user, nil
}

func (r *UserRepository) Create(user User) (UserResponse, error) {
	hashedPassword, err := r.hashPassword(user.Password)
	if err != nil {
		return UserResponse{}, err
	}

	err = r.DB.QueryRow(
		CreateUserQuery,
		user.FirstName, user.LastName, user.Email, user.Age, hashedPassword,
	).Scan(&user.ID)
	if err != nil {
		return UserResponse{}, err
	}
	return UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Age:       user.Age,
	}, nil
}
