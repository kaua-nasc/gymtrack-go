package domain

import (
	"errors"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailNotFound      = errors.New("email not found")
	ErrInvalidCode        = errors.New("invalid code")
	ErrTrainerNotFound    = errors.New("trainer not found")
	ErrAlreadyVerified    = errors.New("email already verified")
	ErrNotVerified        = errors.New("user not verified")
)
