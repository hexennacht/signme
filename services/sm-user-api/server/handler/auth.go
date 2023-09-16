package handler

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/hexennacht/signme/services/sm-user-api/core/entity"
	"github.com/hexennacht/signme/services/sm-user-api/core/services"
	"github.com/hexennacht/signme/services/sm-user-api/grpc/v1/auth"
)

type AuthHandler struct {
	svc services.AuthService

	auth.UnimplementedAuthenticationServer
}

func NewAuthHandler(service services.AuthService) *AuthHandler {
	return &AuthHandler{svc: service}
}

func (a *AuthHandler) SignIn(ctx context.Context, request *auth.SignInRequest) (*auth.SignInResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthHandler) SignUp(ctx context.Context, request *auth.SignUpRequest) (*auth.SignUpResponse, error) {
	resp, err := a.svc.SignUp(ctx, &entity.CreateUser{
		Name:                 request.Name,
		Username:             request.Email,
		Password:             request.Password,
		PasswordConfirmation: request.PasswordConfirmation,
	})
	if err != nil {
		return nil, errors.FromError(err)
	}

	return resp, nil
}

func (a *AuthHandler) SignOut(ctx context.Context, request *auth.SignOutRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}
