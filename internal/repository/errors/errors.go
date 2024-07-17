package errors

import "errors"

var (
	ErrFlightNotFound = errors.New("flight not found")
	ErrUserNotFound = errors.New("user not found")
)
