package user

import (
	"context"
	"time"

	"github.com/hexennacht/signme/services/sm-user-api/core/entity"
	"github.com/hexennacht/signme/services/sm-user-api/ent"
	"github.com/hexennacht/signme/services/sm-user-api/ent/user"
	"github.com/hexennacht/signme/services/sm-user-api/repository"
)

type repo struct {
	db *ent.UserClient
}

func NewRepository(db *ent.UserClient) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) CreateNewUser(ctx context.Context, req *entity.CreateUser) (*entity.User, error) {
	userResult, err := r.db.Create().
		SetName(req.Name).
		SetUsername(req.Username).
		SetPassword(req.Password).
		SetStatus(user.StatusNew).
		SetDeletedAt(time.Time{}).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return userEntityMapper(userResult), nil
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	userResult, err := r.db.Query().
		Where(user.UsernameEqualFold(email)).
		First(ctx)
	if err != nil {
		return nil, err
	}

	return userEntityMapper(userResult), nil
}
