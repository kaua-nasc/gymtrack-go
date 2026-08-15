package subscription

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/kaua-nasc/gymtrack-go/apps/api/training-plan/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
	"github.com/kaua-nasc/gymtrack-go/libs/utils"
)

type Service struct {
	repo     Repository
	identity domain.IdentityClient
}

func NewService(
	repo Repository,
	identity domain.IdentityClient,
) *Service {
	return &Service{
		repo:     repo,
		identity: identity,
	}
}

func (s *Service) ListSubscription(ctx context.Context, userId string, filters domain.ListSubscriptionFilters) ([]*domain.PlanSubscription, error) {
	slog.InfoContext(ctx, "listing subscriptions", slog.String("user_id", userId))

	subscriptions, err := s.repo.ListSubscription(ctx, userId, filters)
	if err != nil {
		slog.ErrorContext(ctx, "failed to list subscriptions", slog.String("user_id", userId), slog.Any("error", err))
		return nil, err
	}

	token, _ := ctx.Value(string(auth.TokenContextKey)).(string)
	user, err := s.identity.FindUser(ctx, userId, token)
	if err != nil {
		slog.WarnContext(ctx, "failed to fetch user details for filtering subscriptions", slog.String("user_id", userId), slog.Any("error", err))
		return subscriptions, nil
	}

	filteredSubs := make([]*domain.PlanSubscription, 0)
	for _, sub := range subscriptions {
		// If PROTECTED and (COMPLETED or CANCELED)
		if sub.TrainingPlan != nil && sub.TrainingPlan.Visibility == domain.Protected &&
			(sub.Status == domain.Completed || sub.Status == domain.Canceled) {

			// Check if Author is current trainer
			if user.StudentOf == nil || user.StudentOf.TrainerId != sub.TrainingPlan.AuthorId {
				continue // Hide it
			}
		}
		filteredSubs = append(filteredSubs, sub)
	}

	return filteredSubs, nil
}

func (s *Service) ListSubscriptionByUserId(ctx context.Context, id, userId string, filters domain.ListSubscriptionFilters) ([]*domain.PlanSubscription, error) {
	slog.InfoContext(ctx, "listing subscriptions", slog.String("user_id", userId))

	token, _ := ctx.Value(string(auth.TokenContextKey)).(string)

	user, err := s.identity.FindUser(ctx, userId, token)
	if err != nil {
		return nil, err
	}

	subscriptions, err := s.repo.ListSubscription(ctx, userId, filters)
	if err != nil {
		slog.ErrorContext(ctx, "failed to list subscriptions", slog.String("user_id", userId), slog.Any("error", err))
		return nil, err
	}

	if userId == id {
		return subscriptions, nil
	}

	if user.StudentOf != nil && user.StudentOf.ID == id {
		privacy, err := s.identity.GetStudentPrivacy(ctx, userId, token)
		if err == nil && privacy != nil && !privacy.ShareTrainingProgress {
			return nil, domain.ErrPrivacySettingsForbidden
		}

		subs := make([]*domain.PlanSubscription, 0)
		for _, s := range subscriptions {
			if s.Type == domain.PrivateSubscription {
				continue
			}
			subs = append(subs, s)
		}
		return subs, nil
	}

	subs := make([]*domain.PlanSubscription, 0)
	for _, s := range subscriptions {
		if s.Type == domain.PrivateSubscription || s.Type == domain.PartialAccessSubscription {
			continue
		}
		subs = append(subs, s)
	}

	return subs, nil
}

func (s *Service) Subscribe(ctx context.Context, planId, userId string) error {
	slog.InfoContext(ctx, "subscribing user to plan", slog.String("plan_id", planId), slog.String("user_id", userId))

	plan, err := s.repo.FindPlanForSubscription(ctx, planId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to fetch plan for subscription", slog.String("plan_id", planId), slog.String("user_id", userId), slog.Any("error", err))
		return err
	}
	if plan == nil {
		slog.WarnContext(ctx, "plan not found for subscription", slog.String("plan_id", planId), slog.String("user_id", userId))
		return domain.ErrPlanNotFound
	}

	if plan.AuthorId == userId {
		slog.WarnContext(ctx, "user attempted to subscribe to own plan", slog.String("plan_id", planId), slog.String("user_id", userId))
		return domain.ErrCannotSubscribeOwnPlan
	}

	alreadySubscribed, isComplete, err := s.repo.GetSubscriptionEligibility(ctx, planId, userId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to check subscription eligibility", slog.String("plan_id", planId), slog.String("user_id", userId), slog.Any("error", err))
		return err
	}

	if alreadySubscribed {
		slog.WarnContext(ctx, "user already subscribed to plan", slog.String("plan_id", planId), slog.String("user_id", userId))
		return domain.ErrAlreadySubscribed
	}

	if !isComplete {
		slog.WarnContext(ctx, "attempted to subscribe to incomplete plan", slog.String("plan_id", planId))
		return domain.ErrPlanIncomplete
	}

	if err := s.authorizeSubscription(ctx, plan, userId); err != nil {
		return err
	}

	if plan.MaxSubscriptions != nil {
		count, err := s.repo.CountActiveSubscriptionsByPlan(ctx, planId)
		if err != nil {
			slog.ErrorContext(ctx, "failed to count active subscriptions", slog.String("plan_id", planId), slog.Any("error", err))
			return err
		}
		if count >= *plan.MaxSubscriptions {
			slog.WarnContext(ctx, "plan reached max subscriptions", slog.String("plan_id", planId), slog.Int("max", *plan.MaxSubscriptions), slog.Int("active", count))
			return domain.ErrMaxSubscriptionsReached
		}
	}

	id, err := utils.GenerateUUIDV7(ctx)
	if err != nil {
		return err
	}
	now := time.Now().UTC()
	sub := &domain.PlanSubscription{
		Id:             *id,
		TrainingPlanId: planId,
		UserId:         userId,
		Status:         domain.NotStarted,
		Type:           domain.TotalAccessSubscription,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	if err := s.repo.CreatePlanSubscription(ctx, sub); err != nil {
		slog.ErrorContext(ctx, "failed to create plan subscription", slog.String("plan_id", planId), slog.String("user_id", userId), slog.Any("error", err))
		return err
	}
	return nil
}

func (s *Service) authorizeSubscription(ctx context.Context, plan *domain.TrainingPlan, userId string) error {
	switch plan.Visibility {
	case domain.Public:
		return nil
	case domain.Protected:
		token, _ := ctx.Value(string(auth.TokenContextKey)).(string)
		user, err := s.identity.FindUser(ctx, userId, token)
		if err != nil {
			slog.WarnContext(ctx, "failed to validate student relation for protected plan", slog.String("plan_id", *plan.Id), slog.String("user_id", userId), slog.Any("error", err))
			return domain.ErrSubscriptionForbidden
		}
		if user.StudentOf == nil || user.StudentOf.TrainerId != plan.AuthorId {
			slog.WarnContext(ctx, "user is not a student of the plan author", slog.String("plan_id", *plan.Id), slog.String("user_id", userId))
			return domain.ErrSubscriptionForbidden
		}
		return nil
	case domain.Private:
		slog.WarnContext(ctx, "attempted to subscribe to private plan", slog.String("plan_id", *plan.Id), slog.String("user_id", userId))
		return domain.ErrSubscriptionForbidden
	}
	return domain.ErrSubscriptionForbidden
}

func (s *Service) Unsubscribe(ctx context.Context, planId, userId string) error {
	slog.InfoContext(ctx, "unsubscribing user from plan", slog.String("plan_id", planId), slog.String("user_id", userId))

	existing, err := s.repo.FindSubscriptionByPlan(ctx, planId, userId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to find subscription for unsubscription", slog.String("plan_id", planId), slog.String("user_id", userId), slog.Any("error", err))
		return err
	}
	if existing == nil {
		slog.WarnContext(ctx, "subscription not found for unsubscription", slog.String("plan_id", planId), slog.String("user_id", userId))
		return errors.New("subscription not found")
	}

	if err := s.repo.UpdateSubscriptionStatus(ctx, *existing, domain.Canceled); err != nil {
		slog.ErrorContext(ctx, "failed to update plan subscription status to canceled", slog.String("plan_id", planId), slog.String("user_id", userId), slog.Any("error", err))
		return err
	}

	return nil
}

func (s *Service) ChangeSubscriptionStatus(ctx context.Context, planId, userId string, status domain.PlanSubscriptionStatus) error {
	subscription, err := s.repo.FindSubscriptionByPlan(ctx, planId, userId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to find subscription for status change", slog.String("plan_id", planId), slog.String("user_id", userId), slog.Any("error", err))
		return err
	}
	if subscription == nil {
		return errors.New("subscription not found")
	}

	switch status {
	case domain.NotStarted:
		if subscription.Status != domain.Canceled {
			return fmt.Errorf("cannot change status from %s to %s", subscription.Status, status)
		}
	case domain.InProgress:
		if subscription.Status != domain.NotStarted {
			return fmt.Errorf("cannot change status from %s to %s", subscription.Status, status)
		}
		inProgressSub, err := s.repo.FindInProgressSubscription(ctx, userId)
		if err != nil {
			return err
		}
		if inProgressSub != nil && inProgressSub.Id != subscription.Id {
			return fmt.Errorf("o usuário já possui uma inscrição em progresso")
		}
	case domain.Completed:
		if subscription.Status != domain.InProgress {
			return fmt.Errorf("cannot change status from %s to %s", subscription.Status, status)
		}
	case domain.Canceled:
		if subscription.Status != domain.InProgress {
			return fmt.Errorf("cannot change status from %s to %s", subscription.Status, status)
		}
	}

	return s.repo.UpdateSubscriptionStatus(ctx, *subscription, status)
}

func (s *Service) ChangeSubscriptionPrivacy(ctx context.Context, planId, userId string, subsType domain.PlanSubscriptionType) error {
	subscription, err := s.repo.FindSubscriptionByPlan(ctx, planId, userId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to find subscription for privacy change", slog.String("plan_id", planId), slog.String("user_id", userId), slog.Any("error", err))
		return err
	}
	if subscription == nil {
		return errors.New("subscription not found")
	}

	return s.repo.UpdateSubscriptionPrivacy(ctx, *subscription, subsType)
}

func (s *Service) CompleteDay(ctx context.Context, subsId, userId, dayId string) error {
	slog.InfoContext(ctx, "completing plan day", slog.String("subscription_id", subsId), slog.String("user_id", userId), slog.String("day_id", dayId))

	sub, err := s.repo.FindSubscription(ctx, subsId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to find subscription for day completion", slog.String("subscription_id", subsId), slog.String("user_id", userId), slog.Any("error", err))
		return err
	}
	if sub == nil {
		slog.WarnContext(ctx, "subscription not found for day completion", slog.String("subscription_id", subsId), slog.String("user_id", userId))
		return errors.New("subscription not found")
	}

	if sub.Status == domain.Completed || sub.Status == domain.Canceled {
		return fmt.Errorf("não é possível modificar o progresso de uma inscrição concluída ou cancelada")
	}

	if sub.Status == domain.NotStarted {
		if err := s.ChangeSubscriptionStatus(ctx, sub.TrainingPlanId, sub.UserId, domain.InProgress); err != nil {
			return err
		}
	}

	progress, err := s.repo.FindInProgressDayProgress(ctx, sub.Id, dayId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to find subscription progress for day completion", slog.Any("error", err))
		return err
	}

	if progress == nil {
		return errors.New("treino não iniciado para este dia")
	}

	if err := s.repo.UpdateSubscriptionProgressStatus(ctx, progress.Id, domain.DayCompleted); err != nil {
		slog.ErrorContext(ctx, "failed to update subscription progress", slog.Any("error", err))
		return err
	}

	progressQuantity, err := s.repo.CountSubscriptionProgress(ctx, sub.Id)
	if err == nil && progressQuantity == sub.TrainingPlan.TimeInDays {
		s.ChangeSubscriptionStatus(ctx, sub.TrainingPlanId, sub.UserId, domain.Completed)
	}

	return nil
}

func (s *Service) CancelDay(ctx context.Context, subsId, userId, dayId string) error {
	slog.InfoContext(ctx, "canceling plan day", slog.String("subscription_id", subsId), slog.String("user_id", userId), slog.String("day_id", dayId))

	sub, err := s.repo.FindSubscription(ctx, subsId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to find subscription for day completion", slog.String("subscription_id", subsId), slog.String("user_id", userId), slog.Any("error", err))
		return err
	}
	if sub == nil {
		slog.WarnContext(ctx, "subscription not found for day completion", slog.String("subscription_id", subsId), slog.String("user_id", userId))
		return errors.New("subscription not found")
	}

	if sub.Status == domain.Completed || sub.Status == domain.Canceled {
		return fmt.Errorf("não é possível modificar o progresso de uma inscrição concluída ou cancelada")
	}

	if sub.Status == domain.NotStarted {
		if err := s.ChangeSubscriptionStatus(ctx, sub.TrainingPlanId, sub.UserId, domain.InProgress); err != nil {
			return err
		}
	}

	progress, err := s.repo.FindInProgressDayProgress(ctx, sub.Id, dayId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to find subscription progress for day completion", slog.Any("error", err))
		return err
	}

	if progress == nil {
		return errors.New("treino não iniciado para este dia")
	}

	if err := s.repo.UpdateSubscriptionProgressStatus(ctx, progress.Id, domain.DayCanceled); err != nil {
		slog.ErrorContext(ctx, "failed to update subscription progress", slog.Any("error", err))
		return err
	}

	progressQuantity, err := s.repo.CountSubscriptionProgress(ctx, sub.Id)
	if err == nil && progressQuantity == sub.TrainingPlan.TimeInDays {
		s.ChangeSubscriptionStatus(ctx, sub.TrainingPlanId, sub.UserId, domain.Completed)
	}

	return nil
}

func (s *Service) StartDay(ctx context.Context, subsId, userId, dayId string) error {
	slog.InfoContext(ctx, "starting plan day", slog.String("subscription_id", subsId), slog.String("user_id", userId), slog.String("day_id", dayId))

	sub, err := s.repo.FindSubscription(ctx, subsId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to find subscription for day start", slog.String("subscription_id", subsId), slog.String("user_id", userId), slog.Any("error", err))
		return err
	}
	if sub == nil {
		slog.WarnContext(ctx, "subscription not found for day start", slog.String("subscription_id", subsId), slog.String("user_id", userId))
		return errors.New("subscription not found")
	}

	if sub.Status == domain.Completed || sub.Status == domain.Canceled {
		return fmt.Errorf("não é possível modificar o progresso de uma inscrição concluída ou cancelada")
	}

	if sub.Status == domain.NotStarted {
		if err := s.ChangeSubscriptionStatus(ctx, sub.TrainingPlanId, sub.UserId, domain.InProgress); err != nil {
			return err
		}
	}

	progress, err := s.repo.FindInProgressDayProgress(ctx, sub.Id, dayId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to find subscription progress for day start", slog.Any("error", err))
		return err
	}

	if progress != nil {
		if err := s.repo.UpdateSubscriptionProgressStatus(ctx, progress.Id, domain.DayInProgress); err != nil {
			slog.ErrorContext(ctx, "failed to update subscription progress", slog.Any("error", err))
			return err
		}
	} else {
		id, err := utils.GenerateUUIDV7(ctx)
		if err != nil {
			return err
		}
		now := time.Now().UTC()
		newProgress := &domain.PlanDayProgress{
			Id:                 *id,
			DayId:              dayId,
			PlanSubscriptionId: sub.Id,
			Status:             domain.DayInProgress,
			CreatedAt:          now,
			UpdatedAt:          now,
		}

		if err := s.repo.CreateSubscriptionProgress(ctx, newProgress); err != nil {
			slog.ErrorContext(ctx, "failed to create subscription progress", slog.Any("error", err))
			return err
		}
	}

	return nil
}

func (s *Service) ListWeeklyDayProgress(ctx context.Context, userId string) (*domain.WeeklyDayProgress, error) {
	progresses, err := s.repo.ListWeeklyDayProgress(ctx, userId)
	if err != nil {
		return nil, err
	}

	weekly := &domain.WeeklyDayProgress{}
	for i := range progresses {
		p := &progresses[i]
		switch p.UpdatedAt.Weekday() {
		case time.Monday:
			weekly.Mon = p
		case time.Tuesday:
			weekly.Tue = p
		case time.Wednesday:
			weekly.Wed = p
		case time.Thursday:
			weekly.Thu = p
		case time.Friday:
			weekly.Fri = p
		case time.Saturday:
			weekly.Sat = p
		case time.Sunday:
			weekly.Sun = p
		}
	}

	return weekly, nil
}

func (s *Service) FindNextDay(ctx context.Context, userId string) (*domain.PlanDayProgress, error) {
	slog.InfoContext(ctx, "finding next day for user", slog.String("user_id", userId))

	activeSub, err := s.repo.FindActiveSubscription(ctx, userId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to fetch active subscription", slog.String("user_id", userId), slog.Any("error", err))
		return nil, err
	}

	if activeSub == nil {
		slog.InfoContext(ctx, "no active subscription found for user", slog.String("user_id", userId))
		return nil, nil
	}

	completedCount, err := s.repo.CountSubscriptionProgress(ctx, activeSub.Id)
	if err != nil {
		slog.ErrorContext(ctx, "failed to count subscription progress", slog.String("subscription_id", activeSub.Id), slog.Any("error", err))
		return nil, err
	}

	if completedCount >= activeSub.TrainingPlan.TimeInDays {
		slog.InfoContext(ctx, "subscription completed by progress count", slog.String("subscription_id", activeSub.Id))
		_ = s.ChangeSubscriptionStatus(ctx, activeSub.TrainingPlanId, activeSub.UserId, domain.Completed)
		return nil, nil
	}

	lastProgress, err := s.repo.FindLastDayProgressByUser(ctx, userId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to fetch last day progress", slog.String("user_id", userId), slog.Any("error", err))
		return nil, err
	}

	var targetDay *domain.Day

	if lastProgress == nil {
		// First day of training - fetch the first day of the plan (lowest sequence)
		targetDay, err = s.repo.FindFirstDay(ctx, activeSub.TrainingPlanId)
		if err != nil {
			slog.ErrorContext(ctx, "failed to fetch first day of training plan", slog.String("plan_id", activeSub.TrainingPlanId), slog.Any("error", err))
			return nil, err
		}
		if targetDay == nil {
			slog.WarnContext(ctx, "training plan has no days", slog.String("plan_id", activeSub.TrainingPlanId))
			return nil, nil
		}
	} else {
		// Progress exists - get the completed day details
		currentDay, err := s.repo.FindDayWithExercises(ctx, lastProgress.DayId)
		if err != nil {
			slog.ErrorContext(ctx, "failed to fetch current day details", slog.String("day_id", lastProgress.DayId), slog.Any("error", err))
			return nil, err
		}
		if currentDay == nil {
			// Fall back to first day of the training plan if the last progress day was deleted
			targetDay, err = s.repo.FindFirstDay(ctx, activeSub.TrainingPlanId)
		} else {
			// Find the next day in sequence
			targetDay, err = s.repo.FindNextDayInSequence(ctx, activeSub.TrainingPlanId, currentDay.Sequence)
			if err != nil {
				slog.ErrorContext(ctx, "failed to fetch next day in sequence", slog.String("plan_id", activeSub.TrainingPlanId), slog.Any("error", err))
				return nil, err
			}
			if targetDay == nil {
				// Circular wrap-around: fall back to the first day of the training plan
				targetDay, err = s.repo.FindFirstDay(ctx, activeSub.TrainingPlanId)
				if err != nil {
					slog.ErrorContext(ctx, "failed to fetch first day during circular wrap-around", slog.String("plan_id", activeSub.TrainingPlanId), slog.Any("error", err))
					return nil, err
				}
			}
		}
	}

	if targetDay == nil {
		return nil, nil
	}

	return &domain.PlanDayProgress{
		DayId:              targetDay.Id,
		Day:                targetDay,
		PlanSubscriptionId: activeSub.Id,
		Status:             domain.DayInProgress,
	}, nil
}

func (s *Service) GetEngagementSummary(ctx context.Context, userId string) (*domain.EngagementSummary, error) {
	return s.repo.GetEngagementSummary(ctx, userId)
}
