package errs

import (
	"errors"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrUserExists         = errors.New("user already exists")
	ErrInvalidFormat      = errors.New("invalid request format")
	ErrRegistration       = errors.New("registration failed")
	ErrInvalidToken       = errors.New("invalid or expired token")
	ErrUserIDNotFound     = errors.New("user_id not found in token")
	ErrNotFound           = errors.New("not found")
	ErrInvalidArgument    = errors.New("invalid argument")
	ErrConflict           = errors.New("conflict")
)
