package domain

import "errors"

var (
	ErrPrivacySettingsForbidden = errors.New("student privacy settings do not allow this action")
)
