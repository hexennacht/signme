package entity

import "time"

type CreateUser struct {
	Name                 string `validate:"required"`
	Username             string `validate:"email"`
	Password             string `validate:"required,alphanum,gt=6"`
	PasswordConfirmation string `validate:"required,eqfield=Password"`
}

type UserStatus string

type User struct {
	ID             int64
	Username       string
	Password       string
	ProfilePicture string
	Status         UserStatus
	CreatedAt      time.Time
	UpdatedAt      *time.Time
	DeletedAt      *time.Time
}
