package entity

import "time"

type CreateUser struct {
	Name                 string `validate:"required"`
	Username             string `validate:"required,email"`
	Password             string `validate:"required,alphanum,gt=6"`
	PasswordConfirmation string `validate:"required,eqfield=password"`
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
