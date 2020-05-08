package exceptions

import (
	"errors"
)

var (
	ErrUserNotFound    = errors.New("User Not Found")
	ErrUserEmailExists = errors.New("User Email Exists")
	ErrGeneral         = errors.New("General Error")
)
