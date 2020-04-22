package exceptions

import (
	"errors"
)

var (
	ErrUserNotFound = errors.New("User Not Found")
)
