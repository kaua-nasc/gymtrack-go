package internal

import (
	"time"
)

// Enums
type UserType string

const (
	Admin   UserType = "ADMIN"
	Trainer UserType = "TRAINER"
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
	ID                string     `json:"id" validate:"required,uuid4"`
	CreatedAt         time.Time  `json:"createdAt"`
	UpdatedAt         time.Time  `json:"updatedAt"`
	FirstName         string     `json:"firstName" validate:"required,min=1,max=255"`
	LastName          string     `json:"lastName" validate:"required,min=1,max=255"`
	Email             string     `json:"email" validate:"required,email"`
	Bio               *string    `json:"bio,omitempty"`
	ProfilePictureUrl *string    `json:"profilePictureUrl,omitempty"`
	Password          string     `json:"password" validate:"required,min=8"`
	Type              UserType   `json:"type" validate:"required"`
	Height            *float64   `json:"height,omitempty"`
	CurrentWeight     *float64   `json:"currentWeight,omitempty"`
	WeightUnit        WeightUnit `json:"weightUnit" validate:"required"`
	HeightUnit        HeightUnit `json:"heightUnit" validate:"required"`
	TrainerInviteCode *string    `json:"trainerInviteCode,omitempty"`
	Cref              *string    `json:"cref,omitempty"`
	IsVerified        bool       `json:"isVerified"`

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
	ID            string              `json:"id" validate:"required,uuid4"`
	CreatedAt     time.Time           `json:"createdAt"`
	UpdatedAt     time.Time           `json:"updatedAt"`
	Type          BodyMeasurementType `json:"type" validate:"required"`
	Value         float64             `json:"value" validate:"required"`
	MeasuredAt    time.Time           `json:"measuredAt" validate:"required"`
	UserId        string              `json:"userId" validate:"required,uuid4"`
	TrainerNote   *string             `json:"trainerNote,omitempty"`
	TrainerNoteAt *time.Time          `json:"trainerNoteAt,omitempty"`

	// Relations
	User *User `json:"user,omitempty"`
}

// MetricGoal Entity
type MetricGoal struct {
	ID            string           `json:"id" validate:"required,uuid4"`
	CreatedAt     time.Time        `json:"createdAt"`
	UpdatedAt     time.Time        `json:"updatedAt"`
	Type          string           `json:"type" validate:"required"`
	StartingValue float64          `json:"startingValue" validate:"required"`
	TargetValue   float64          `json:"targetValue" validate:"required"`
	Deadline      *time.Time       `json:"deadline,omitempty"`
	AchievedAt    *time.Time       `json:"achievedAt,omitempty"`
	Status        MetricGoalStatus `json:"status" validate:"required"`
	UserId        string           `json:"userId" validate:"required,uuid4"`

	// Relations
	User *User `json:"user,omitempty"`
}

// TrainerStudentRelation Entity
type TrainerStudentRelation struct {
	ID        string    `json:"id" validate:"required,uuid4"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	TrainerId string    `json:"trainerId" validate:"required,uuid4"`
	StudentId string    `json:"studentId" validate:"required,uuid4"`
	LinkedAt  time.Time `json:"linkedAt" validate:"required"`

	// Relations
	Trainer *User `json:"trainer,omitempty"`
	Student *User `json:"student,omitempty"`
}

// UserFollows Entity
type UserFollows struct {
	ID          string    `json:"id" validate:"required,uuid4"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	FollowerId  string    `json:"followerId" validate:"required,uuid4"`
	FollowingId string    `json:"followingId" validate:"required,uuid4"`

	// Relations
	Follower  *User `json:"follower,omitempty"`
	Following *User `json:"following,omitempty"`
}

// UserPrivacySettings Entity
type UserPrivacySettings struct {
	ID                       string    `json:"id" validate:"required,uuid4"`
	CreatedAt                time.Time `json:"createdAt"`
	UpdatedAt                time.Time `json:"updatedAt"`
	ShareName                bool      `json:"shareName"`
	ShareEmail               bool      `json:"shareEmail"`
	ShareTrainingProgress    bool      `json:"shareTrainingProgress"`
	SharePastDataWithTrainer bool      `json:"sharePastDataWithTrainer"`
	UserId                   string    `json:"userId" validate:"required,uuid4"`

	// Relations
	User *User `json:"user,omitempty"`
}

// WeightLog Entity
type WeightLog struct {
	ID            string     `json:"id" validate:"required,uuid4"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	Weight        float64    `json:"weight" validate:"required"`
	MeasuredAt    time.Time  `json:"measuredAt" validate:"required"`
	UserId        string     `json:"userId" validate:"required,uuid4"`
	TrainerNote   *string    `json:"trainerNote,omitempty"`
	TrainerNoteAt *time.Time `json:"trainerNoteAt,omitempty"`

	// Relations
	User *User `json:"user,omitempty"`
}
