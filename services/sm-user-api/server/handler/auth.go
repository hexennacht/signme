package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/hexennacht/signme/services/sm-user-api/grpc/v1/auth"
)

type AuthHandler struct {
	svc interface{}

	auth.UnimplementedAuthenticationServer
}

func NewAuthHandler(service interface{}) *AuthHandler {
	return &AuthHandler{svc: service}
}

func (a *AuthHandler) SignIn(ctx context.Context, request *auth.SignInRequest) (*auth.SignInResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthHandler) SignUp(ctx context.Context, request *auth.SignUpRequest) (*auth.SignUpResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthHandler) SignOut(ctx context.Context, request *auth.SignOutRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}
