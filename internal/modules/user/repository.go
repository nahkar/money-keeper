package user

import (
	"database/sql"
	"fmt"

	"dario.cat/mergo"
	"github.com/nahkar/money-keeper/internal/utils"
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

func (r *UserRepository) FindAll() ([]User, error) {
	rows, err := r.DB.Query(FindAllUsersQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Age); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) FindById(id int) (User, error) {
	var user User

	err := r.DB.QueryRow(FindUserByIDQuery, id).
		Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Age, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, fmt.Errorf("%w: user with id %d", ErrUserNotFound, id)
		}
		return User{}, err
	}

	return user, nil
}

func (r *UserRepository) Create(user CreateUserRequest) (User, error) {
	hashedPassword, err := r.hashPassword(user.Password)
	if err != nil {
		return User{}, err
	}
	result, err := r.DB.Exec(
		CreateUserQuery,
		user.FirstName, user.LastName, user.Email, user.Age, hashedPassword,
	)
	if err != nil {
		return User{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return User{}, err
	}

	return User{
		ID:        int(id),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Age:       user.Age,
	}, nil
}

func (r *UserRepository) Update(id int, user UpdateUserRequest) (User, error) {
	currentUser, err := r.FindById(id)
	if err != nil {
		return User{}, err
	}

	partialUpdate := User{
		FirstName: utils.DerefString(user.FirstName),
		LastName:  utils.DerefString(user.LastName),
		Email:     utils.DerefString(user.Email),
		Age:       utils.DerefInt(user.Age),
		Password:  utils.DerefString(user.Password),
	}

	if user.Password != nil {
		hashedPassword, err := r.hashPassword(*user.Password)
		if err != nil {
			return User{}, err
		}
		partialUpdate.Password = hashedPassword
	}

	if err := mergo.Merge(&currentUser, partialUpdate, mergo.WithOverride); err != nil {
		return User{}, err
	}

	result, err := r.DB.Exec(
		UpdateUserQuery,
		currentUser.FirstName, currentUser.LastName, currentUser.Email,
		currentUser.Age, currentUser.Password, id,
	)
	if err != nil {
		return User{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return User{}, err
	}
	if rowsAffected == 0 {
		return User{}, fmt.Errorf("%w: user with id %d", ErrUserNotFound, id)
	}

	return currentUser, nil
}

func (r *UserRepository) Delete(id int) error {
	result, err := r.DB.Exec(DeleteUserQuery, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%w: user with id %d", ErrUserNotFound, id)
	}

	return nil
}
