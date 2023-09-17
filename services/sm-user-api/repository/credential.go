package repository

import (
	"context"

	"github.com/hexennacht/signme/services/sm-user-api/core/entity"
)

type CredentialRepository interface {
	InsertCredentials(ctx context.Context, userID int64, userCredential *entity.UserCredential) error
}
