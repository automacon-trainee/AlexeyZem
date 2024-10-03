package custom_error

import (
	"errors"
)

var (
	ErrBadUser       = errors.New("invalid username or password")
	ErrAlreadyExists = errors.New("user already exists")
)
