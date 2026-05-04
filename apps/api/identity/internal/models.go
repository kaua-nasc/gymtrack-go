package internal

import (
	"time"
)

type UserType string

const (
	Admin   UserType = "admin"
	Trainer UserType = "trainer"
	Client  UserType = "client"
)

type User struct {
	ID        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Type      UserType  `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserFollows struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	FollowerId  string    `json:"followerId"`
	FollowingId string    `json:"followingId"`
}
