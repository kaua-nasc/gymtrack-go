package domain

import (
	"errors"
)

var (
	ErrUserAlreadyExists         = errors.New("user already exists")
	ErrInvalidCredentials        = errors.New("invalid credentials")
	ErrUserNotFound              = errors.New("user not found")
	ErrEmailNotFound             = errors.New("email not found")
	ErrInvalidCode               = errors.New("invalid code")
	ErrTrainerNotFound           = errors.New("trainer not found")
	ErrAlreadyVerified           = errors.New("email already verified")
	ErrNotVerified               = errors.New("user not verified")
	ErrUnauthorizedAccess        = errors.New("unauthorized access to this data")
	ErrUnauthorizedTrainerAccess = errors.New("trainer does not have an active link with this student")
	ErrPrivacySettingsForbidden  = errors.New("student privacy settings do not allow this action")
	ErrInvalidPeriod             = errors.New("invalid period: maximum 1 year allowed")
	ErrForbiddenRole             = errors.New("access denied for your user type")
)
