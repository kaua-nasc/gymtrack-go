package domain

import "errors"

var (
	ErrPrivacySettingsForbidden = errors.New("student privacy settings do not allow this action")
	ErrFeedbackNotAllowed       = errors.New("only subscribed users can provide feedback for this plan")
	ErrFeedbackNotFound         = errors.New("feedback not found")
	ErrNotFeedbackAuthor        = errors.New("only the author can delete this feedback")
	ErrPlanNotFound             = errors.New("training plan not found")
	ErrAlreadySubscribed        = errors.New("already subscribed to this training plan")
	ErrPlanIncomplete           = errors.New("training plan is incomplete (must have at least one day and one exercise)")
	ErrCannotSubscribeOwnPlan   = errors.New("you cannot subscribe to your own training plan")
	ErrSubscriptionForbidden    = errors.New("you are not allowed to subscribe to this training plan")
	ErrMaxSubscriptionsReached  = errors.New("training plan reached the maximum number of subscriptions")
)
