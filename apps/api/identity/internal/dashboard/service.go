package dashboard

import (
	"context"
	"time"

	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/domain"
	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/metrics"
	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/trainer"
	"github.com/kaua-nasc/gymtrack-go/apps/api/identity/internal/user"
	"github.com/kaua-nasc/gymtrack-go/libs/auth"
)

type Service struct {
	userRepo           user.Repository
	trainerRepo        trainer.Repository
	metricsRepo        metrics.Repository
	trainingPlanClient domain.TrainingPlanClient
}

func NewService(userRepo user.Repository, trainerRepo trainer.Repository, metricsRepo metrics.Repository, trainingPlanClient domain.TrainingPlanClient) *Service {
	return &Service{
		userRepo:           userRepo,
		trainerRepo:        trainerRepo,
		metricsRepo:        metricsRepo,
		trainingPlanClient: trainingPlanClient,
	}
}

func (s *Service) GetStudentEngagement(ctx context.Context, trainerId, studentId string) (*domain.EngagementSummary, error) {
	// 1. Check if they are linked
	linkedAt, err := s.trainerRepo.GetTrainerLinkDate(ctx, trainerId, studentId)
	if err != nil {
		return nil, err
	}
	if linkedAt == nil {
		return nil, domain.ErrUnauthorizedTrainerAccess
	}

	// 2. Check privacy settings
	privacy, err := s.userRepo.GetPrivacySettings(ctx, studentId)
	if err != nil {
		return nil, err
	}

	if privacy != nil && !privacy.ShareTrainingProgress {
		return nil, domain.ErrPrivacySettingsForbidden
	}

	// 3. Fetch summary from training-plan microservice
	token, _ := ctx.Value(string(auth.TokenContextKey)).(string)
	summary, err := s.trainingPlanClient.GetEngagementSummary(ctx, studentId, token)
	if err != nil {
		return nil, err
	}

	return summary, nil
}

func (s *Service) GetStudentBiometrics(ctx context.Context, trainerId, studentId string, start, end time.Time) (*domain.BiometricDashboard, error) {
	// 1. Validation: Max 1 year
	if end.Sub(start).Hours() > 24*365 {
		return nil, domain.ErrInvalidPeriod
	}

	// 2. Check Link
	linkedAt, err := s.trainerRepo.GetTrainerLinkDate(ctx, trainerId, studentId)
	if err != nil {
		return nil, err
	}
	if linkedAt == nil {
		return nil, domain.ErrUnauthorizedTrainerAccess
	}

	// 3. Check Privacy
	privacy, err := s.userRepo.GetPrivacySettings(ctx, studentId)
	if err != nil {
		return nil, err
	}

	dashboard := &domain.BiometricDashboard{
		Measurements: make(map[string]*domain.BiometricMeasurementSummary),
		Charts: domain.BiometricCharts{
			Weight:       make([]domain.ChartPoint, 0),
			Measurements: make(map[string][]domain.ChartPoint),
		},
	}

	u, _ := s.userRepo.Find(ctx, studentId, studentId)

	// 4. Weight Data
	if privacy == nil || privacy.ShareWeightLogs {
		history, _ := s.metricsRepo.ListWeightHistory(ctx, studentId, start, end)
		if len(history) > 0 {
			current := history[len(history)-1]
			summary := &domain.BiometricWeightSummary{
				Current: current.Weight,
			}

			// Calculate deltas
			summary.Delta7Days = calculateDelta(history, current.Weight, 7)
			summary.Delta30Days = calculateDelta(history, current.Weight, 30)
			summary.Delta90Days = calculateDelta(history, current.Weight, 90)

			dashboard.Weight = summary

			// Chart
			for _, p := range history {
				dashboard.Charts.Weight = append(dashboard.Charts.Weight, domain.ChartPoint{
					Date:  p.MeasuredAt,
					Value: p.Weight,
				})
			}

			// BMI
			if u != nil && u.Height != nil && *u.Height > 0 {
				h := *u.Height / 100 // cm to m
				bmi := current.Weight / (h * h)
				dashboard.CurrentBMI = &bmi
			}
		}
	}

	// 5. Measurement Data
	if privacy == nil || privacy.ShareBodyMeasurements {
		mHistory, _ := s.metricsRepo.ListMeasurementsHistory(ctx, studentId, start, end)

		// Group by type
		byType := make(map[string][]*domain.BodyMeasurement)
		for _, m := range mHistory {
			byType[string(m.Type)] = append(byType[string(m.Type)], m)
		}

		for mType, items := range byType {
			if len(items) > 0 {
				current := items[len(items)-1]
				summary := &domain.BiometricMeasurementSummary{
					Current: current.Value,
				}
				if len(items) > 1 {
					summary.Delta = current.Value - items[len(items)-2].Value
				}
				dashboard.Measurements[mType] = summary

				// Chart
				chartPoints := make([]domain.ChartPoint, 0)
				for _, p := range items {
					chartPoints = append(chartPoints, domain.ChartPoint{
						Date:  p.MeasuredAt,
						Value: p.Value,
					})
				}
				dashboard.Charts.Measurements[mType] = chartPoints
			}
		}
	}

	return dashboard, nil
}

func calculateDelta(history []*domain.WeightLog, currentVal float64, days int) float64 {
	targetDate := time.Now().UTC().AddDate(0, 0, -days)
	var closest *domain.WeightLog
	for _, l := range history {
		if l.MeasuredAt.Before(targetDate) || l.MeasuredAt.Equal(targetDate) {
			closest = l
		} else {
			break
		}
	}
	if closest != nil {
		return currentVal - closest.Weight
	}
	return 0
}
