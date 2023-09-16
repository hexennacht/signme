package user

import (
	"github.com/hexennacht/signme/services/sm-user-api/core/entity"
	"github.com/hexennacht/signme/services/sm-user-api/ent"
)

func userEntityMapper(userResult *ent.User) *entity.User {
	return &entity.User{
		Username:       userResult.Username,
		Password:       userResult.Password,
		ProfilePicture: userResult.ProfilePicture,
		Status:         entity.UserStatus(userResult.Status.String()),
		CreatedAt:      userResult.CreatedAt,
		UpdatedAt:      userResult.UpdatedAt,
		DeletedAt:      userResult.DeletedAt,
	}
}
