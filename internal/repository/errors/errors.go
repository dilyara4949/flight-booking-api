package errors

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrInvalidEmailPassword = errors.New("invalid email or password")
)
