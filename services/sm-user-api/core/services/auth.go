package services

import (
	"context"

	"github.com/hexennacht/signme/services/sm-user-api/grpc/v1/auth"
)

type AuthService interface {
	SignIn(ctx context.Context, req *auth.SignInRequest) (*auth.SignInResponse, error)
	SignUp(ctx context.Context, req *auth.SignUpRequest) (*auth.SignUpResponse, error)
	SignOut(ctx context.Context, req *auth.SignOutRequest) error
}

type authService struct{}

func NewAuthService() AuthService {
	return &authService{}
}

func (a *authService) SignIn(ctx context.Context, req *auth.SignInRequest) (*auth.SignInResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *authService) SignUp(ctx context.Context, req *auth.SignUpRequest) (*auth.SignUpResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *authService) SignOut(ctx context.Context, req *auth.SignOutRequest) error {
	//TODO implement me
	panic("implement me")
}
