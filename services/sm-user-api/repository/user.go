package repository

import (
	"context"

	"github.com/hexennacht/signme/services/sm-user-api/core/entity"
)

type UserRepository interface {
	CreateNewUser(ctx context.Context, user *entity.CreateUser) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
}
