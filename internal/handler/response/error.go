package response

import "errors"

type Error struct {
	Error string `json:"error"`
}

var ErrEmptyRequestFields = errors.New("request fields cannot be empty")
