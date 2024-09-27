package myerror

import (
	"errors"
)

var (
	ErrNotBook           = errors.New("not book")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrWrongPassword     = errors.New("wrong password or email")
	ErrUserNotFound      = errors.New("user not found")
	ErrWrongToken        = errors.New("wrong token")
)
