package internal

import (
	"errors"
)

var (
	BadRequestError error = errors.New("bad request")
	DeletedError    error = errors.New("user deleted")
)
