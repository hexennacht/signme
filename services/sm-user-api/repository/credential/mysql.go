package credential

import (
	"context"
	"time"

	"github.com/hexennacht/signme/services/sm-user-api/core/entity"
	"github.com/hexennacht/signme/services/sm-user-api/ent"
	"github.com/hexennacht/signme/services/sm-user-api/ent/credential"
	"github.com/hexennacht/signme/services/sm-user-api/repository"
)

type repo struct {
	db *ent.CredentialClient
}

func NewRepository(db *ent.CredentialClient) repository.CredentialRepository {
	return &repo{db: db}
}

func (r *repo) InsertCredentials(ctx context.Context, userID int64, userCredential *entity.UserCredential) error {
	private := r.db.Create().
		SetCredential(*userCredential.Private).
		SetCredentialType(credential.CredentialTypePrivate).
		SetDeletedAt(time.Time{}).
		AddUserIDs(userID)

	public := r.db.Create().
		SetCredential(*userCredential.Public).
		SetCredentialType(credential.CredentialTypePublic).
		SetDeletedAt(time.Time{}).
		AddUserIDs(userID)

	return r.db.CreateBulk(private, public).Exec(ctx)
}
