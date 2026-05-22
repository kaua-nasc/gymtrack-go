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
	identity *domain.IdentityService
}

func NewService(
	repo Repository,
	identity *domain.IdentityService,
) *Service {
	return &Service{
		repo:     repo,
		identity: identity,
	}
}

func (s *Service) ListSubscription(ctx context.Context, userId string) ([]*domain.PlanSubscription, error) {
	slog.InfoContext(ctx, "listing subscriptions", slog.String("user_id", userId))

	subscriptions, err := s.repo.ListSubscription(ctx, userId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to list subscriptions", slog.String("user_id", userId), slog.Any("error", err))
		return nil, err
	}

	return subscriptions, nil
}

func (s *Service) ListSubscriptionByUserId(ctx context.Context, id, userId string) ([]*domain.PlanSubscription, error) {
	slog.InfoContext(ctx, "listing subscriptions", slog.String("user_id", userId))

	token, _ := ctx.Value(string(auth.TokenContextKey)).(string)

	user, err := s.identity.FindUser(ctx, userId, token)
	if err != nil {
		return nil, err
	}

	subscriptions, err := s.repo.ListSubscription(ctx, userId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to list subscriptions", slog.String("user_id", userId), slog.Any("error", err))
		return nil, err
	}

	if userId == id {
		return subscriptions, nil
	}

	if user.StudentOf != nil && user.StudentOf.ID == id {
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

func (s *Service) Subscribe(ctx context.Context, planId, userId string, subType domain.PlanSubscriptionType) error {
	slog.InfoContext(ctx, "subscribing user to plan", slog.String("plan_id", planId), slog.String("user_id", userId))

	alreadySubscribed, isComplete, err := s.repo.GetSubscriptionEligibility(ctx, planId, userId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to check subscription eligibility", slog.String("plan_id", planId), slog.String("user_id", userId), slog.Any("error", err))
		return err
	}

	if alreadySubscribed {
		slog.WarnContext(ctx, "user already subscribed to plan", slog.String("plan_id", planId), slog.String("user_id", userId))
		return errors.New("already subscribed")
	}

	if !isComplete {
		slog.WarnContext(ctx, "attempted to subscribe to incomplete plan", slog.String("plan_id", planId))
		return errors.New("training plan is incomplete (must have at least one day and one exercise)")
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
		Type:           subType,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	if err := s.repo.CreatePlanSubscription(ctx, sub); err != nil {
		slog.ErrorContext(ctx, "failed to create plan subscription", slog.String("plan_id", planId), slog.String("user_id", userId), slog.Any("error", err))
		return err
	}
	return nil
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

	if err := s.repo.DeletePlanSubscription(ctx, existing); err != nil {
		slog.ErrorContext(ctx, "failed to delete plan subscription", slog.String("plan_id", planId), slog.String("user_id", userId), slog.Any("error", err))
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
			return fmt.Errorf("o status deve ser cancelado")
		}
	case domain.InProgress:
		if subscription.Status != domain.NotStarted {
			return fmt.Errorf("o status deve ser cancelado")
		}
	case domain.Completed:
		if subscription.Status != domain.InProgress {
			return fmt.Errorf("o status deve ser cancelado")
		}
	case domain.Canceled:
		if subscription.Status != domain.InProgress {
			return fmt.Errorf("o status deve ser cancelado")
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

	id, err := utils.GenerateUUIDV7(ctx)
	if err != nil {
		return err
	}
	now := time.Now().UTC()
	progress := &domain.PlanDayProgress{
		Id:                 *id,
		DayId:              dayId,
		PlanSubscriptionId: sub.Id,
		Status:             domain.DayCompleted,
		CreatedAt:          now,
		UpdatedAt:          now,
	}

	if err := s.repo.CreateSubscriptionProgress(ctx, progress); err != nil {
		slog.ErrorContext(ctx, "failed to create subscription progress", slog.Any("error", err))
		return err
	}

	progressQuantity, err := s.repo.CountSubscriptionProgress(ctx, sub.Id)
	if err == nil && progressQuantity == sub.TrainingPlan.TimeInDays {
		s.ChangeSubscriptionStatus(ctx, sub.TrainingPlanId, sub.UserId, domain.Completed)
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
