package domain

import "time"

// Enums
type PostEntityType string

const (
	TrainingPlanPost PostEntityType = "TRAINING_PLAN"
)

type PostStatus string

const (
	PostPending  PostStatus = "PENDING"
	PostApproved PostStatus = "APPROVED"
	PostRejected PostStatus = "REJECTED"
)

type Post struct {
	Id             *string         `json:"id" validate:"required,uuid"`
	CreatedAt      time.Time       `json:"createdAt"`
	UpdatedAt      time.Time       `json:"updatedAt"`
	AuthorId       string          `json:"-" validate:"required,uuid"`
	Content        string          `json:"content" validate:"required"`
	MediaUrls      []string        `json:"mediaUrls"`
	EntityId       *string         `json:"entityId" validate:"omitempty,uuid"`
	EntityType     *PostEntityType `json:"entityType" validate:"omitempty"`
	Status         PostStatus      `json:"status"`
	RejectedReason *string         `json:"rejectedReason,omitempty"`

	// Relations
	Likes    []Like    `json:"likes,omitempty"`
	Comments []Comment `json:"comments,omitempty"`

	// Virtual fields
	LikesCount         int  `json:"likesCount"`
	CommentsCount      int  `json:"commentsCount"`
	LikedByCurrentUser bool `json:"likedByCurrentUser"`
	Author             any  `json:"author,omitempty"`
	Entity             any  `json:"trainingPlan,omitempty"`
}

type Like struct {
	Id        string    `json:"id" validate:"required,uuid"`
	CreatedAt time.Time `json:"createdAt"`
	UserId    string    `json:"userId" validate:"required,uuid"`
	PostId    string    `json:"postId" validate:"required,uuid"`
}

type Comment struct {
	Id        *string   `json:"id" validate:"required,uuid"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Content   string    `json:"content" validate:"required"`
	AuthorId  string    `json:"authorId" validate:"required,uuid"`
	PostId    string    `json:"postId" validate:"required,uuid"`

	// Virtual fields
	Author any `json:"author,omitempty"`
}

type AuditLog struct {
	Id             string     `json:"id"`
	PostId         string     `json:"postId"`
	AdminId        string     `json:"adminId"`
	PreviousStatus PostStatus `json:"previousStatus"`
	NewStatus      PostStatus `json:"newStatus"`
	Reason         *string    `json:"reason,omitempty"`
	CreatedAt      time.Time  `json:"createdAt"`

	// Relations
	Post   *Post `json:"post,omitempty"`
	Author any   `json:"author,omitempty"`
	Admin  any   `json:"admin,omitempty"`
}
