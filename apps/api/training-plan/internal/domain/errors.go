package domain

import "errors"

var (
	ErrPrivacySettingsForbidden = errors.New("student privacy settings do not allow this action")
	ErrFeedbackNotAllowed       = errors.New("only subscribed users can provide feedback for this plan")
)
