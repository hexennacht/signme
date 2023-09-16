package user

import (
	"context"

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
		SetUsername(req.Username).
		SetPassword(req.Password).
		SetStatus(user.StatusNew).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &entity.User{
		Username:       userResult.Username,
		Password:       userResult.Password,
		ProfilePicture: userResult.ProfilePicture,
		Status:         entity.UserStatus(userResult.Status.String()),
		CreatedAt:      userResult.CreatedAt,
		UpdatedAt:      userResult.UpdatedAt,
		DeletedAt:      userResult.DeletedAt,
	}, nil
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}
