package cmd

import (
	"errors"
)

var (
	ErrConnection      = errors.New("Connection error")
	ErrNotFound        = errors.New("Not found")
	ErrInvalidResponse = errors.New("Invalid server response")
	ErrNotNumber       = errors.New("Not a number")
)
