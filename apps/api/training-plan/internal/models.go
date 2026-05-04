package internal

import "time"

type CursorData struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}

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
	DaySkipped    PlanDayProgressStatus = "SKIPPED"
)

// TrainingPlan Entity
type TrainingPlan struct {
	Id               string                 `json:"id" validate:"required,uuid4"`
	CreatedAt        time.Time              `json:"createdAt"`
	UpdatedAt        time.Time              `json:"updatedAt"`
	Name             string                 `json:"name" validate:"required,min=3,max=255"`
	AuthorId         string                 `json:"authorId" validate:"required,uuid4"`
	TimeInDays       int                    `json:"timeInDays" validate:"required,min=1"`
	Type             TrainingPlanType       `json:"type" validate:"required"`
	Visibility       TrainingPlanVisibility `json:"visibility" validate:"required"`
	Level            TrainingPlanLevel      `json:"level" validate:"required"`
	Observation      *string                `json:"observation,omitempty"`
	Pathology        *string                `json:"pathology,omitempty"`
	MaxSubscriptions *int                   `json:"maxSubscriptions,omitempty"`
	ImageUrl         *string                `json:"imageUrl"`
	Description      *string                `json:"description"`

	// Relations
	Days                []Day                  `json:"days,omitempty"`
	PlanSubscriptions   []PlanSubscription     `json:"planSubscriptions,omitempty"`
	AccessRequests      []PlanAccessRequest    `json:"accessRequests,omitempty"`
	PrivateParticipants []PlanParticipant      `json:"privateParticipants,omitempty"`
	Feedbacks           []TrainingPlanFeedback `json:"feedbacks,omitempty"`
	Likes               []TrainingPlanLike     `json:"likes,omitempty"`
	Comments            []TrainingPlanComment  `json:"comments,omitempty"`
	Invites             []PlanInvite           `json:"invites,omitempty"`

	// Computed/Virtual Fields
	LikesCount             int    `json:"likesCount"`
	LikedByCurrentUser     bool   `json:"likedByCurrentUser"`
	PlanSubscriptionStatus string `json:"planSubscriptionStatus"`
	Author                 any    `json:"author,omitempty"`
}

// Day Entity
type Day struct {
	Id             string    `json:"id" validate:"required,uuid4"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Name           string    `json:"name" validate:"required,min=1,max=255"`
	TrainingPlanId string    `json:"trainingPlanId" validate:"required,uuid4"`

	// Relations
	TrainingPlan    *TrainingPlan     `json:"trainingPlan,omitempty"`
	Exercises       []Exercise        `json:"exercises,omitempty"`
	PlanDayProgress []PlanDayProgress `json:"planDayProgress,omitempty"`
}

// Exercise Entity
type Exercise struct {
	Id          string       `json:"id" validate:"required,uuid4"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	Name        string       `json:"name" validate:"required,min=1,max=255"`
	DayId       string       `json:"dayId" validate:"required,uuid4"`
	Type        ExerciseType `json:"type" validate:"required"`
	SetsNumber  int          `json:"setsNumber" validate:"required,min=1"`
	RepsNumber  int          `json:"repsNumber" validate:"required,min=1"`
	Description *string      `json:"description,omitempty"`
	Observation *string      `json:"observation,omitempty"`

	// Relations
	Day *Day `json:"day,omitempty"`
}

// PlanSubscription Entity
type PlanSubscription struct {
	Id             string                 `json:"id" validate:"required,uuid4"`
	CreatedAt      time.Time              `json:"createdAt"`
	UpdatedAt      time.Time              `json:"updatedAt"`
	TrainingPlanId string                 `json:"trainingPlanId" validate:"required,uuid4"`
	UserId         string                 `json:"userId" validate:"required,uuid4"`
	Status         PlanSubscriptionStatus `json:"status" validate:"required"`
	Type           PlanSubscriptionType   `json:"type" validate:"required"`

	// Relations
	PlanDayProgress []PlanDayProgress                `json:"planDayProgress,omitempty"`
	TrainingPlan    *TrainingPlan                    `json:"trainingPlan,omitempty"`
	PrivacySettings *PlanSubscriptionPrivacySettings `json:"privacySettings,omitempty"`
}

// PlanAccessRequest Entity
type PlanAccessRequest struct {
	Id             string                  `json:"id" validate:"required,uuid4"`
	CreatedAt      time.Time               `json:"createdAt"`
	UpdatedAt      time.Time               `json:"updatedAt"`
	UserId         string                  `json:"userId" validate:"required,uuid4"`
	TrainingPlanId string                  `json:"trainingPlanId" validate:"required,uuid4"`
	Status         PlanAccessRequestStatus `json:"status" validate:"required"`

	// Relations
	TrainingPlan *TrainingPlan `json:"trainingPlan,omitempty"`
}

// Support Entities
// PlanSubscriptionPrivacySettings Entity
type PlanSubscriptionPrivacySettings struct {
	Id                   string    `json:"id" validate:"required,uuid4"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
	PlanSubscriptionId   string    `json:"planSubscriptionId" validate:"required,uuid4"`
	ShareProgress        bool      `json:"shareProgress"`
	SharePersonalMetrics bool      `json:"sharePersonalMetrics"`

	// Relations
	PlanSubscription *PlanSubscription `json:"planSubscription,omitempty"`
}

// PlanParticipant Entity
type PlanParticipant struct {
	Id             string    `json:"id" validate:"required,uuid4"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	UserId         string    `json:"userId" validate:"required,uuid4"`
	TrainingPlanId string    `json:"trainingPlanId" validate:"required,uuid4"`
	ExpirationDate time.Time `json:"expirationDate"`
	ApprovedAt     time.Time `json:"approvedAt"`

	// Relations
	TrainingPlan *TrainingPlan `json:"trainingPlan,omitempty"`
}

// TrainingPlanFeedback Entity
type TrainingPlanFeedback struct {
	Id             string    `json:"id" validate:"required,uuid4"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	TrainingPlanId string    `json:"trainingPlanId" validate:"required,uuid4"`
	UserId         string    `json:"userId" validate:"required,uuid4"`
	Rating         float64   `json:"rating" validate:"required,min=0,max=5"`
	Message        *string   `json:"message,omitempty"`

	// Relations
	TrainingPlan *TrainingPlan `json:"trainingPlan,omitempty"`
}

// TrainingPlanLike Entity
type TrainingPlanLike struct {
	Id             string    `json:"id" validate:"required,uuid4"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	LikedBy        string    `json:"likedBy" validate:"required,uuid4"`
	TrainingPlanId string    `json:"trainingPlanId" validate:"required,uuid4"`

	// Relations
	TrainingPlan *TrainingPlan `json:"trainingPlan,omitempty"`
}

// TrainingPlanComment Entity
type TrainingPlanComment struct {
	Id             string    `json:"id" validate:"required,uuid4"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Content        string    `json:"content" validate:"required"`
	AuthorId       string    `json:"authorId" validate:"required,uuid4"`
	TrainingPlanId string    `json:"trainingPlanId" validate:"required,uuid4"`

	// Relations
	TrainingPlan *TrainingPlan `json:"trainingPlan,omitempty"`

	// Virtual fields
	Author any `json:"author,omitempty"`
}

// PlanInvite Entity
type PlanInvite struct {
	Id             string           `json:"id" validate:"required,uuid4"`
	CreatedAt      time.Time        `json:"createdAt"`
	UpdatedAt      time.Time        `json:"updatedAt"`
	PlanId         string           `json:"planId" validate:"required,uuid4"`
	SenderId       string           `json:"senderId" validate:"required,uuid4"`
	RecipientId    *string          `json:"recipientId,omitempty" validate:"omitempty,uuid4"`
	RecipientEmail string           `json:"recipientEmail" validate:"required,email"`
	Status         PlanInviteStatus `json:"status" validate:"required"`

	// Relations
	TrainingPlan *TrainingPlan `json:"trainingPlan,omitempty"`
}

// ExerciseLog Entity
type ExerciseLog struct {
	Id         string    `json:"id" validate:"required,uuid4"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	UserId     string    `json:"userId" validate:"required,uuid4"`
	ExerciseId string    `json:"exerciseId" validate:"required,uuid4"`
	Reps       []int     `json:"reps" validate:"required"`
	Weight     []float64 `json:"weight" validate:"required"`
	Notes      *string   `json:"notes,omitempty"`

	// Relations
	Exercise *Exercise `json:"exercise,omitempty"`
}

type PlanDayProgress struct {
	Id                 string                `json:"id" validate:"required,uuid4"`
	CreatedAt          time.Time             `json:"createdAt"`
	UpdatedAt          time.Time             `json:"updatedAt"`
	DayId              string                `json:"dayId" validate:"required,uuid4"`
	PlanSubscriptionId string                `json:"planSubscriptionId" validate:"required,uuid4"`
	Status             PlanDayProgressStatus `json:"status" validate:"required"`

	// Relations
	Day              *Day              `json:"day,omitempty"`
	PlanSubscription *PlanSubscription `json:"planSubscription,omitempty"`
}

