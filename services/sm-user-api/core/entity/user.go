package entity

import "time"

type CreateUser struct {
	Username             string
	Password             string
	PasswordConfirmation string
}

type UserStatus string

type User struct {
	Username       string
	Password       string
	ProfilePicture string
	Status         UserStatus
	CreatedAt      time.Time
	UpdatedAt      *time.Time
	DeletedAt      *time.Time
}
