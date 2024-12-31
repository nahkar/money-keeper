package user

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidID    = errors.New("invalid user id")
)
