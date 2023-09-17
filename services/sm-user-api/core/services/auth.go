package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"

	sshutils "github.com/hexennacht/signme/share-libs/ssh-utils"

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
	userRepo       repository.UserRepository
	credentialRepo repository.CredentialRepository
}

func NewAuthService(userRepo repository.UserRepository, credentialRepo repository.CredentialRepository) AuthService {
	return &authService{userRepo: userRepo, credentialRepo: credentialRepo}
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

	user, err := a.userRepo.GetUserByEmail(ctx, req.Username)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if user != nil {
		message := fmt.Sprintf("User with email %s is already exists", user.Username)
		return nil, errors.BadRequest(auth.AuthenticationError_USER_EXISTS.String(), message)
	}

	// encrypt user password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.PasswordConfirmation), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.InternalServer(auth.AuthenticationError_INTERNAL_SERVER_ERROR.String(), err.Error())
	}

	req.Password = string(hashedPassword)

	// save user data to database
	result, err := a.userRepo.CreateNewUser(ctx, req)
	if err != nil {
		return nil, errors.InternalServer(auth.AuthenticationError_INTERNAL_SERVER_ERROR.String(), err.Error())
	}

	// generate ssh for registered user
	ssh := sshutils.NewGenerateSSH(strings.Split(req.Username, "@")[0], req.Password)

	if err := ssh.GenerateSSHKey(); err != nil {
		return nil, errors.InternalServer(auth.AuthenticationError_INTERNAL_SERVER_ERROR.String(), err.Error())
	}

	// store public & private user ssh to database
	publicName, publicContent := ssh.GetPublicKey()
	privateName, privateContent := ssh.GetPrivateKey()

	userCredential := &entity.UserCredential{
		Private: &entity.Credential{
			Name:    privateName,
			Content: privateContent,
		},
		Public: &entity.Credential{
			Name:    publicName,
			Content: publicContent,
		},
	}

	err = a.credentialRepo.InsertCredentials(ctx, result.ID, userCredential)
	if err != nil {
		return nil, errors.InternalServer(auth.AuthenticationError_INTERNAL_SERVER_ERROR.String(), err.Error())
	}

	return &auth.SignUpResponse{
		Email:        result.Username,
		Token:        "",
		RefreshToken: "",
	}, nil
}

func (a *authService) SignOut(ctx context.Context, req *auth.SignOutRequest) error {
	//TODO implement me
	panic("implement me")
}
