package internal

import (
	"time"
)

type UserType string

const (
	Admin  UserType = "admin"
	Trainer UserType = "trainer"
	Client  UserType = "client"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Type      UserType  `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
