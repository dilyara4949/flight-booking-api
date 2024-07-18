package errors

import "errors"

var (
	ErrTicketNotFound = errors.New("ticket not found")
	ErrUserNotFound   = errors.New("user not found")
)
