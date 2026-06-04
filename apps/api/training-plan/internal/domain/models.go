package domain

import "time"

// Enums
type TrainingPlanType string

const (
	Hypertrophy TrainingPlanType = "HYPERTROPHY"
	Strength    TrainingPlanType = "STRENGTH"
	Mixed       TrainingPlanType = "MIXED"
)

type TrainingPlanVisibility string

const (
	Public    TrainingPlanVisibility = "PUBLIC"
	Protected TrainingPlanVisibility = "PROTECTED"
	Private   TrainingPlanVisibility = "PRIVATE"
)

type TrainingPlanLevel string

const (
	Beginner     TrainingPlanLevel = "BEGINNER"
	Intermediate TrainingPlanLevel = "INTERMEDIATE"
	Advanced     TrainingPlanLevel = "ADVANCED"
)

type ExerciseType string

const (
	Warmup      ExerciseType = "WARMUP"
	Recognition ExerciseType = "RECOGNITION"
	Work        ExerciseType = "WORK"
	Cardio      ExerciseType = "CARDIO"
)

type PlanSubscriptionStatus string

const (
	NotStarted PlanSubscriptionStatus = "NOT_STARTED"
	InProgress PlanSubscriptionStatus = "IN_PROGRESS"
	Completed  PlanSubscriptionStatus = "COMPLETED"
	Canceled   PlanSubscriptionStatus = "CANCELED"
)

type PlanSubscriptionType string

const (
	TotalAccessSubscription   PlanSubscriptionType = "TOTAL_ACCESS"
	PartialAccessSubscription PlanSubscriptionType = "PARTIAL_ACCESS"
	PrivateSubscription       PlanSubscriptionType = "PRIVATE"
)

type PlanAccessRequestStatus string

const (
	PendingAccess  PlanAccessRequestStatus = "PENDING"
	ApprovedAccess PlanAccessRequestStatus = "APPROVED"
	RejectedAccess PlanAccessRequestStatus = "REJECTED"
	CanceledAccess PlanAccessRequestStatus = "CANCELED"
)

type PlanInviteStatus string

const (
	PendingInvite  PlanInviteStatus = "PENDING"
	AcceptedInvite PlanInviteStatus = "ACCEPTED"
	RejectedInvite PlanInviteStatus = "REJECTED"
	CanceledInvite PlanInviteStatus = "CANCELED"
)

type PlanDayProgressStatus string

const (
	DayInProgress PlanDayProgressStatus = "IN_PROGRESS"
	DayCompleted  PlanDayProgressStatus = "COMPLETED"
	DayCanceled   PlanDayProgressStatus = "CANCELLED"
)

type ListSubscriptionFilters struct {
	Status     *PlanSubscriptionStatus
	Type       *PlanSubscriptionType
	PlanType   *TrainingPlanType
	Visibility *TrainingPlanVisibility
	Level      *TrainingPlanLevel
	AuthorId   *string
}

type WeeklyDayProgress struct {
	Mon *PlanDayProgress `json:"mon"`
	Tue *PlanDayProgress `json:"tue"`
	Wed *PlanDayProgress `json:"wed"`
	Thu *PlanDayProgress `json:"thu"`
	Fri *PlanDayProgress `json:"fri"`
	Sat *PlanDayProgress `json:"sat"`
	Sun *PlanDayProgress `json:"sun"`
}

// TrainingPlan Entity
type TrainingPlan struct {
	Id                *string                `json:"id,omitempty" validate:"omitempty,uuid"`
	CreatedAt         time.Time              `json:"createdAt"`
	UpdatedAt         time.Time              `json:"updatedAt"`
	Name              string                 `json:"name" validate:"required,min=3,max=255"`
	AuthorId          string                 `json:"authorId" validate:"required,uuid"`
	TimeInDays        int                    `json:"timeInDays" validate:"required,min=1"`
	Type              TrainingPlanType       `json:"type" validate:"required"`
	Visibility        TrainingPlanVisibility `json:"visibility" validate:"required"`
	Level             TrainingPlanLevel      `json:"level" validate:"required"`
	Observation       *string                `json:"observation,omitempty"`
	Pathology         *string                `json:"pathology,omitempty"`
	MaxSubscriptions  *int                   `json:"maxSubscriptions,omitempty"`
	ImageUrl          *string                `json:"imageUrl"`
	Description       *string                `json:"description"`
	TotalRatingSum    *float64               `json:"totalRatingSum"`
	TotalRatingsCount int                    `json:"totalRatingsCount"`

	// Relations
	Days                []Day                  `json:"days,omitempty"`
	PlanSubscriptions   []PlanSubscription     `json:"planSubscriptions,omitempty"`
	AccessRequests      []PlanAccessRequest    `json:"accessRequests,omitempty"`
	PrivateParticipants []PlanParticipant      `json:"privateParticipants,omitempty"`
	Feedbacks           []TrainingPlanFeedback `json:"feedbacks,omitempty"`
	Invites             []PlanInvite           `json:"invites,omitempty"`

	// Computed/Virtual Fields
	PlanSubscriptionStatus *PlanSubscriptionStatus `json:"planSubscriptionStatus"`
	Author                 any                     `json:"author,omitempty"`
}

// Day Entity
type Day struct {
	Id             string    `json:"id" validate:"required,uuid"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Name           string    `json:"name" validate:"required,min=1,max=255"`
	TrainingPlanId string    `json:"trainingPlanId" validate:"required,uuid"`
	Sequence       int       `json:"sequence"`

	// Relations
	TrainingPlan    *TrainingPlan     `json:"trainingPlan,omitempty"`
	Exercises       []Exercise        `json:"exercises,omitempty"`
	PlanDayProgress []PlanDayProgress `json:"planDayProgress,omitempty"`
}

// Exercise Entity
type Exercise struct {
	Id          string       `json:"id" validate:"required,uuid" form:"id"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	Name        string       `json:"name" validate:"required,min=1,max=255" form:"name"`
	DayId       string       `json:"dayId" validate:"required,uuid" form:"dayId"`
	Type        ExerciseType `json:"type" validate:"required" form:"type"`
	SetsNumber  int          `json:"setsNumber" validate:"required,min=1" form:"setsNumber"`
	RepsNumber  int          `json:"repsNumber" validate:"required,min=1" form:"repsNumber"`
	Description *string      `json:"description,omitempty" form:"description"`
	Observation *string      `json:"observation,omitempty" form:"observation"`
	Sequence    int          `json:"sequence" form:"sequence"`
	VideoUrl    *string      `json:"videoUrl,omitempty" form:"videoUrl"`
	ImageUrl    *string      `json:"imageUrl,omitempty" form:"imageUrl"`

	// Relations
	Day *Day `json:"day,omitempty"`
}

// PlanSubscription Entity
type PlanSubscription struct {
	Id                 string                 `json:"id" validate:"required,uuid"`
	CreatedAt          time.Time              `json:"createdAt"`
	UpdatedAt          time.Time              `json:"updatedAt"`
	TrainingPlanId     string                 `json:"trainingPlanId" validate:"required,uuid"`
	UserId             string                 `json:"userId" validate:"required,uuid"`
	Status             PlanSubscriptionStatus `json:"status" validate:"required"`
	Type               PlanSubscriptionType   `json:"type" validate:"required"`
	CompletedDaysCount *int                   `json:"completedDaysCount,omitempty"`

	// Relations
	PlanDayProgress []PlanDayProgress `json:"planDayProgress,omitempty"`
	TrainingPlan    *TrainingPlan     `json:"trainingPlan,omitempty"`
}

// PlanAccessRequest Entity
type PlanAccessRequest struct {
	Id             string                  `json:"id" validate:"required,uuid"`
	CreatedAt      time.Time               `json:"createdAt"`
	UpdatedAt      time.Time               `json:"updatedAt"`
	UserId         string                  `json:"userId" validate:"required,uuid"`
	TrainingPlanId string                  `json:"trainingPlanId" validate:"required,uuid"`
	Status         PlanAccessRequestStatus `json:"status" validate:"required"`

	// Relations
	TrainingPlan *TrainingPlan `json:"trainingPlan,omitempty"`
}

// Support Entities
// PlanParticipant Entity
type PlanParticipant struct {
	Id             string    `json:"id" validate:"required,uuid"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	UserId         string    `json:"userId" validate:"required,uuid"`
	TrainingPlanId string    `json:"trainingPlanId" validate:"required,uuid"`
	ExpirationDate time.Time `json:"expirationDate"`
	ApprovedAt     time.Time `json:"approvedAt"`

	// Relations
	TrainingPlan *TrainingPlan `json:"trainingPlan,omitempty"`
}

// TrainingPlanFeedback Entity
type TrainingPlanFeedback struct {
	Id             string    `json:"id" validate:"required,uuid"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	TrainingPlanId string    `json:"trainingPlanId" validate:"required,uuid"`
	UserId         string    `json:"userId" validate:"required,uuid"`
	Rating         float64   `json:"rating" validate:"required,min=0,max=5"`
	Message        *string   `json:"message,omitempty"`

	// Relations
	TrainingPlan *TrainingPlan `json:"trainingPlan,omitempty"`
}

// PlanInvite Entity
type PlanInvite struct {
	Id             string           `json:"id" validate:"required,uuid"`
	CreatedAt      time.Time        `json:"createdAt"`
	UpdatedAt      time.Time        `json:"updatedAt"`
	PlanId         string           `json:"planId" validate:"required,uuid"`
	SenderId       string           `json:"senderId" validate:"required,uuid"`
	RecipientId    *string          `json:"recipientId,omitempty" validate:"omitempty,uuid"`
	RecipientEmail string           `json:"recipientEmail" validate:"required,email"`
	Status         PlanInviteStatus `json:"status" validate:"required"`

	// Relations
	TrainingPlan *TrainingPlan `json:"trainingPlan,omitempty"`
}

// ExerciseLog Entity
type ExerciseLog struct {
	Id         string    `json:"id" validate:"required,uuid"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	UserId     string    `json:"userId" validate:"required,uuid"`
	ExerciseId string    `json:"exerciseId" validate:"required,uuid"`
	Reps       []int     `json:"reps" validate:"required"`
	Weight     []float64 `json:"weight" validate:"required"`
	Notes      *string   `json:"notes,omitempty"`

	// Relations
	Exercise *Exercise `json:"exercise,omitempty"`
}

type PlanDayProgress struct {
	Id                 string                `json:"id" validate:"required,uuid"`
	CreatedAt          time.Time             `json:"createdAt"`
	UpdatedAt          time.Time             `json:"updatedAt"`
	DayId              string                `json:"dayId" validate:"required,uuid"`
	PlanSubscriptionId string                `json:"planSubscriptionId" validate:"required,uuid"`
	Status             PlanDayProgressStatus `json:"status" validate:"required"`

	// Relations
	Day              *Day              `json:"day,omitempty"`
	PlanSubscription *PlanSubscription `json:"planSubscription,omitempty"`
}

type EngagementSummary struct {
	AdherenceRate      float64    `json:"adherenceRate"`      // % of completed days vs total days in plan
	WeeklyFrequency    int        `json:"weeklyFrequency"`    // Number of days completed in current week
	CurrentPlanName    string     `json:"currentPlanName"`    // Name of the active training plan
	PlanProgress       float64    `json:"planProgress"`       // % of plan completed
	LastWorkoutDate    *time.Time `json:"lastWorkoutDate"`    // Date of the last exercise log or session
	ActiveDaysThisWeek []string   `json:"activeDaysThisWeek"` // Days of the week (e.g., ["Monday", "Wednesday"])
}
