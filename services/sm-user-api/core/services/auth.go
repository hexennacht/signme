package services

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"

	"github.com/hexennacht/signme/services/sm-user-api/core/entity"
	"github.com/hexennacht/signme/services/sm-user-api/ent"
	"github.com/hexennacht/signme/services/sm-user-api/grpc/v1/auth"
	"github.com/hexennacht/signme/services/sm-user-api/repository"
)

type AuthService interface {
	SignIn(ctx context.Context, req *auth.SignInRequest) (*auth.SignInResponse, error)
	SignUp(ctx context.Context, req *entity.CreateUser) (*auth.SignUpResponse, error)
	SignOut(ctx context.Context, req *auth.SignOutRequest) error
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{repo: userRepo}
}

func (a *authService) SignIn(ctx context.Context, req *auth.SignInRequest) (*auth.SignInResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *authService) SignUp(ctx context.Context, req *entity.CreateUser) (*auth.SignUpResponse, error) {
	// validate the request
	v := validator.New()

	if err := v.Struct(req); err != nil {
		return nil, err
	}

	user, err := a.repo.GetUserByEmail(ctx, req.Username)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if user != nil {
		message := fmt.Sprintf("User with email %s is already exists", user.Username)
		return nil, errors.BadRequest(auth.AuthenticationError_USER_EXISTS.String(), message)
	}

	// generate ssh for registered user

	// encrypt user password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.PasswordConfirmation), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.InternalServer(auth.AuthenticationError_INTERNAL_SERVER_ERROR.String(), err.Error())
	}

	req.Password = string(hashedPassword)

	// save user data to database
	result, err := a.repo.CreateNewUser(ctx, req)
	if err != nil {
		return nil, errors.InternalServer(auth.AuthenticationError_INTERNAL_SERVER_ERROR.String(), err.Error())
	}

	// store public user ssh to database

	// store user private ssh to storage

	return nil, nil
}

func (a *authService) SignOut(ctx context.Context, req *auth.SignOutRequest) error {
	//TODO implement me
	panic("implement me")
}
