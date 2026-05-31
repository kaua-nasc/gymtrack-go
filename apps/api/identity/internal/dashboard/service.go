package dashboard

import (
	"context"
	"fmt"
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

func (s *Service) GetStudentInsights(ctx context.Context, trainerId, studentId string) (*domain.InsightsDashboard, error) {
	// 1. Check Link
	linkedAt, err := s.trainerRepo.GetTrainerLinkDate(ctx, trainerId, studentId)
	if err != nil {
		return nil, err
	}
	if linkedAt == nil {
		return nil, domain.ErrUnauthorizedTrainerAccess
	}

	// 2. Check Privacy
	privacy, err := s.userRepo.GetPrivacySettings(ctx, studentId)
	if err != nil {
		return nil, err
	}

	insights := []domain.DashboardInsight{}

	// 3. Review Notifications (Weight & Measurements)
	if privacy == nil || privacy.ShareWeightLogs || privacy.ShareBodyMeasurements {
		wCount, mCount, err := s.metricsRepo.CountUnreviewedMetrics(ctx, studentId)
		if err == nil {
			if (privacy == nil || privacy.ShareWeightLogs) && wCount > 0 {
				insights = append(insights, domain.DashboardInsight{
					Type:        "PENDING_REVIEW_WEIGHT",
					Title:       "Novos registros de peso",
					Description: fmt.Sprintf("O aluno possui %d registros de peso aguardando sua revisão.", wCount),
					Severity:    domain.InsightInfo,
				})
			}
			if (privacy == nil || privacy.ShareBodyMeasurements) && mCount > 0 {
				insights = append(insights, domain.DashboardInsight{
					Type:        "PENDING_REVIEW_MEASUREMENTS",
					Title:       "Novas medidas corporais",
					Description: fmt.Sprintf("O aluno possui %d novas medidas aguardando sua revisão.", mCount),
					Severity:    domain.InsightInfo,
				})
			}
		}
	}

	// 4. Inactivity & Performance Alerts (Requires training-plan data)
	if privacy == nil || privacy.ShareTrainingProgress {
		token, _ := ctx.Value(string(auth.TokenContextKey)).(string)
		summary, err := s.trainingPlanClient.GetEngagementSummary(ctx, studentId, token)
		if err == nil {
			// Inactivity Alert (> 3 days)
			if summary.LastWorkoutDate != nil {
				daysInactive := int(time.Since(*summary.LastWorkoutDate).Hours() / 24)
				if daysInactive >= 3 {
					insights = append(insights, domain.DashboardInsight{
						Type:        "INACTIVITY_ALERT",
						Title:       "Inatividade detectada",
						Description: fmt.Sprintf("O aluno não registra treinos há %d dias.", daysInactive),
						Severity:    domain.InsightWarning,
					})
				}
			} else {
				// No workouts at all?
				insights = append(insights, domain.DashboardInsight{
					Type:        "NO_WORKOUTS_ALERT",
					Title:       "Sem registros de treino",
					Description: "O aluno ainda não realizou nenhum treino deste plano.",
					Severity:    domain.InsightInfo,
				})
			}
		}
	}

	// 5. Stagnation Alert (Weight Plateau - last 30 days)
	if privacy == nil || privacy.ShareWeightLogs {
		end := time.Now().UTC()
		start := end.AddDate(0, 0, -30)
		history, _ := s.metricsRepo.ListWeightHistory(ctx, studentId, start, end)
		if len(history) >= 4 { // Need at least some data points over the month
			minW, maxW := history[0].Weight, history[0].Weight
			for _, l := range history {
				if l.Weight < minW {
					minW = l.Weight
				}
				if l.Weight > maxW {
					maxW = l.Weight
				}
			}
			// If variation is less than 0.3kg over 30 days
			if maxW-minW < 0.3 {
				insights = append(insights, domain.DashboardInsight{
					Type:        "STAGNATION_WEIGHT",
					Title:       "Possível estagnação de peso",
					Description: "O peso do aluno apresentou pouca variação nos últimos 30 dias.",
					Severity:    domain.InsightInfo,
				})
			}
		}
	}

	return &domain.InsightsDashboard{Insights: insights}, nil
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
