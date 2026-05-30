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
	Chest      BodyMeasurementType = "CHEST"
	Waist      BodyMeasurementType = "WAIST"
	Hips       BodyMeasurementType = "HIPS"
	ArmLeft    BodyMeasurementType = "ARM_LEFT"
	ArmRight   BodyMeasurementType = "ARM_RIGHT"
	ThighLeft  BodyMeasurementType = "THIGH_LEFT"
	ThighRight BodyMeasurementType = "THIGH_RIGHT"
	CalfLeft   BodyMeasurementType = "CALF_LEFT"
	CalfRight  BodyMeasurementType = "CALF_RIGHT"
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
