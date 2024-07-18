package errors

import "errors"

var (
	ErrTickerNotFound = errors.New("ticket not found")
	ErrUserNotFound   = errors.New("user not found")
)
