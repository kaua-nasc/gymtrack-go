package domain

import (
	"time"
)

// Enums
type UserType string

const (
	Trainer UserType = "PERSONAL_TRAINER"
	Client  UserType = "CLIENT"
)

type WeightUnit string

const (
	KG WeightUnit = "kg"
	LB WeightUnit = "lb"
)

type HeightUnit string

const (
	CM HeightUnit = "cm"
	IN HeightUnit = "in"
)

type BodyMeasurementType string

const (
	Chest           BodyMeasurementType = "CHEST"
	Waist           BodyMeasurementType = "WAIST"
	Hips            BodyMeasurementType = "HIPS"
	Neck            BodyMeasurementType = "NECK"
	Shoulders       BodyMeasurementType = "SHOULDERS"
	BicepLeft       BodyMeasurementType = "BICEP_LEFT"
	BicepRight      BodyMeasurementType = "BICEP_RIGHT"
	ForearmLeft     BodyMeasurementType = "FOREARM_LEFT"
	ForearmRight    BodyMeasurementType = "FOREARM_RIGHT"
	ThighLeft       BodyMeasurementType = "THIGH_LEFT"
	ThighRight      BodyMeasurementType = "THIGH_RIGHT"
	CalfLeft        BodyMeasurementType = "CALF_LEFT"
	CalfRight       BodyMeasurementType = "CALF_RIGHT"
	BodyFat         BodyMeasurementType = "BODY_FAT"
	WaterPercentage BodyMeasurementType = "WATER_PERCENTAGE"
	MuscleMass      BodyMeasurementType = "MUSCLE_MASS"
	BoneMass        BodyMeasurementType = "BONE_MASS"
)

type MetricGoalStatus string

const (
	MetricGoalActive    MetricGoalStatus = "ACTIVE"
	MetricGoalAchieved  MetricGoalStatus = "ACHIEVED"
	MetricGoalAbandoned MetricGoalStatus = "ABANDONED"
)

// User Entity
type User struct {
	ID                *string    `json:"id" validate:"uuid"`
	CreatedAt         time.Time  `json:"createdAt"`
	UpdatedAt         time.Time  `json:"updatedAt"`
	FirstName         string     `json:"firstName" validate:"required,min=1,max=255"`
	LastName          string     `json:"lastName" validate:"required,min=1,max=255"`
	Email             *string    `json:"email" validate:"required,email"`
	Bio               *string    `json:"bio,omitempty"`
	ProfilePictureUrl *string    `json:"profilePictureUrl,omitempty"`
	Password          string     `json:"password" validate:"required,min=8"`
	Type              UserType   `json:"type"`
	Height            *float64   `json:"height,omitempty"`
	CurrentWeight     *float64   `json:"currentWeight,omitempty"`
	WeightUnit        WeightUnit `json:"weightUnit"`
	HeightUnit        HeightUnit `json:"heightUnit"`
	TrainerInviteCode *string    `json:"trainerInviteCode,omitempty"`
	Cref              *string    `json:"cref,omitempty"`
	IsVerified        bool       `json:"isVerified"`
	IsFollowing       *bool      `json:"isFollowing,omitempty"`

	// Relations
	Measurements    []BodyMeasurement        `json:"measurements,omitempty"`
	MetricGoals     []MetricGoal             `json:"metricGoals,omitempty"`
	TrainerOf       []TrainerStudentRelation `json:"trainerOf,omitempty"`
	StudentOf       *TrainerStudentRelation  `json:"studentOf,omitempty"`
	Followers       []UserFollows            `json:"followers,omitempty"`
	Following       []UserFollows            `json:"following,omitempty"`
	PrivacySettings *UserPrivacySettings     `json:"privacySettings,omitempty"`
	WeightLogs      []WeightLog              `json:"weightLogs,omitempty"`
}

// BodyMeasurement Entity
type BodyMeasurement struct {
	ID            string              `json:"id" validate:"required,uuid"`
	CreatedAt     time.Time           `json:"createdAt"`
	UpdatedAt     time.Time           `json:"updatedAt"`
	Type          BodyMeasurementType `json:"type" validate:"required"`
	Value         float64             `json:"value" validate:"required"`
	MeasuredAt    time.Time           `json:"measuredAt" validate:"required"`
	UserId        string              `json:"userId" validate:"required,uuid"`
	TrainerNote   *string             `json:"trainerNote,omitempty"`
	TrainerNoteAt *time.Time          `json:"trainerNoteAt,omitempty"`

	// Relations
	User *User `json:"user,omitempty"`
}

// MetricGoal Entity
type MetricGoal struct {
	ID            string           `json:"id" validate:"required,uuid"`
	CreatedAt     time.Time        `json:"createdAt"`
	UpdatedAt     time.Time        `json:"updatedAt"`
	Type          string           `json:"type" validate:"required"`
	StartingValue float64          `json:"startingValue" validate:"required"`
	TargetValue   float64          `json:"targetValue" validate:"required"`
	Deadline      *time.Time       `json:"deadline,omitempty"`
	AchievedAt    *time.Time       `json:"achievedAt,omitempty"`
	Status        MetricGoalStatus `json:"status" validate:"required"`
	UserId        string           `json:"userId" validate:"required,uuid"`

	// Relations
	User *User `json:"user,omitempty"`
}

// TrainerStudentRelation Entity
type TrainerStudentRelation struct {
	ID        *string   `json:"id" validate:"required,uuid"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	TrainerId *string   `json:"trainerId" validate:"required,uuid"`
	StudentId string    `json:"studentId" validate:"required,uuid"`
	LinkedAt  time.Time `json:"linkedAt" validate:"required"`

	// Relations
	Trainer *User `json:"trainer,omitempty"`
	Student *User `json:"student,omitempty"`
}

// UserFollows Entity
type UserFollows struct {
	ID          *string   `json:"id" validate:"required,uuid"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	FollowerId  *string   `json:"followerId" validate:"required,uuid"`
	FollowingId *string   `json:"followingId" validate:"required,uuid"`

	// Relations
	Follower  *User `json:"follower,omitempty"`
	Following *User `json:"following,omitempty"`
}

// UserPrivacySettings Entity
type UserPrivacySettings struct {
	ID                       string    `json:"id"`
	CreatedAt                time.Time `json:"createdAt"`
	UpdatedAt                time.Time `json:"updatedAt"`
	ShareEmail               bool      `json:"shareEmail"`
	ShareTrainingProgress    bool      `json:"shareTrainingProgress"`
	SharePastDataWithTrainer bool      `json:"sharePastDataWithTrainer"`
	ShareBodyMeasurements    bool      `json:"shareBodyMeasurements"`
	ShareWeightLogs          bool      `json:"shareWeightLogs"`
	ShareMetricGoals         bool      `json:"shareMetricGoals"`
	AllowTrainerNotes        bool      `json:"allowTrainerNotes"`
	UserId                   string    `json:"userId" validate:"required,uuid"`

	// Relations
	User *User `json:"user,omitempty"`
}

// WeightLog Entity
type WeightLog struct {
	ID            string     `json:"id" validate:"required,uuid"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	Weight        float64    `json:"weight" validate:"required"`
	MeasuredAt    time.Time  `json:"measuredAt" validate:"required"`
	UserId        string     `json:"userId" validate:"required,uuid"`
	TrainerNote   *string    `json:"trainerNote,omitempty"`
	TrainerNoteAt *time.Time `json:"trainerNoteAt,omitempty"`

	// Relations
	User *User `json:"user,omitempty"`
}

type EngagementSummary struct {
	AdherenceRate      float64    `json:"adherenceRate"`      // % of completed days vs total days in plan
	WeeklyFrequency    int        `json:"weeklyFrequency"`    // Number of days completed in current week
	CurrentPlanName    string     `json:"currentPlanName"`    // Name of the active training plan
	PlanProgress       float64    `json:"planProgress"`       // % of plan completed
	LastWorkoutDate    *time.Time `json:"lastWorkoutDate"`    // Date of the last exercise log or session
	ActiveDaysThisWeek []string   `json:"activeDaysThisWeek"` // Days of the week (e.g., ["Monday", "Wednesday"])
}

type BiometricDashboard struct {
	CurrentBMI   *float64                                `json:"currentBmi"`
	Weight       *BiometricWeightSummary                 `json:"weight"`
	Measurements map[string]*BiometricMeasurementSummary `json:"measurements"`
	Charts       BiometricCharts                         `json:"charts"`
}

type BiometricWeightSummary struct {
	Current     float64 `json:"current"`
	Delta7Days  float64 `json:"delta7Days"`
	Delta30Days float64 `json:"delta30Days"`
	Delta90Days float64 `json:"delta90Days"`
}

type BiometricMeasurementSummary struct {
	Current float64 `json:"current"`
	Delta   float64 `json:"delta"` // Compared to previous entry in period
}

type BiometricCharts struct {
	Weight       []ChartPoint            `json:"weight"`
	Measurements map[string][]ChartPoint `json:"measurements"`
}

type ChartPoint struct {
	Date  time.Time `json:"date"`
	Value float64   `json:"value"`
}

type InsightSeverity string

const (
	InsightInfo     InsightSeverity = "INFO"
	InsightWarning  InsightSeverity = "WARNING"
	InsightCritical InsightSeverity = "CRITICAL"
)

type DashboardInsight struct {
	Type        string          `json:"type"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Severity    InsightSeverity `json:"severity"`
	Metadata    map[string]any  `json:"metadata,omitempty"`
}

type InsightsDashboard struct {
	Insights []DashboardInsight `json:"insights"`
}
