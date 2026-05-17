package internal

import "time"

// Enums
type PostEntityType string

const (
	TrainingPlanPost PostEntityType = "TRAINING_PLAN"
)

type Post struct {
	Id         *string         `json:"id" validate:"required,uuid"`
	CreatedAt  time.Time       `json:"createdAt"`
	UpdatedAt  time.Time       `json:"updatedAt"`
	AuthorId   string          `json:"-" validate:"required,uuid"`
	Content    string          `json:"content" validate:"required"`
	EntityId   *string         `json:"-" validate:"omitempty,uuid"`
	EntityType *PostEntityType `json:"-" validate:"omitempty"`

	// Relations
	Likes    []Like    `json:"likes,omitempty"`
	Comments []Comment `json:"comments,omitempty"`

	// Virtual fields
	LikesCount         int  `json:"likesCount"`
	CommentsCount      int  `json:"commentsCount"`
	LikedByCurrentUser bool `json:"likedByCurrentUser"`
	Author             any  `json:"author,omitempty"`
	TrainingPlan       any  `json:"trainingPlan,omitempty"`
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
