package errors

import "errors"

var (
	ErrTicketNotFound = errors.New("ticket not found")
	ErrFlightNotFound       = errors.New("flight not found")
	ErrUserNotFound         = errors.New("user not found")
	ErrInvalidEmailPassword = errors.New("invalid email or password")
)
