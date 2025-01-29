package usecase

import (
	"context"

	authEntity "hiyoko-fiber/internal/domain/entities/auth"
	"hiyoko-fiber/internal/domain/entities/users"
	"hiyoko-fiber/internal/domain/services"
	"hiyoko-fiber/internal/pkg/auth/v1"
	entUtil "hiyoko-fiber/internal/pkg/ent/util"
	"hiyoko-fiber/internal/presentation/http/app/input"
	"hiyoko-fiber/internal/shared"
	logger "hiyoko-fiber/pkg/logging/file"
	"hiyoko-fiber/utils"
)

type UserUseCase interface {
	GetUser(ctx context.Context, id entUtil.ULID) (*users.UserEntity, error)
	Signup(ctx context.Context, input *input.SignupInput) (*authEntity.AuthenticationEntity, error)
	Signin(ctx context.Context, input *input.SigninInput) (*authEntity.AuthenticationEntity, error)
}

type userUseCase struct {
	services.UserRepository
}

func NewUserUseCase(r services.UserRepository) UserUseCase {
	return &userUseCase{r}
}

func (u *userUseCase) GetUser(ctx context.Context, id entUtil.ULID) (*users.UserEntity, error) {
	user, err := u.UserRepository.Get(ctx, &id)
	if err != nil {
		logger.Error("Error getting user", "id", id, "error", err)
		return &users.UserEntity{}, err
	}

	return user, nil
}

func (u *userUseCase) Signup(ctx context.Context, input *input.SignupInput) (*authEntity.AuthenticationEntity, error) {
	err := validateForSignup(u.UserRepository, ctx, input.Email, input.OriginalID)
	if err != nil {
		return &authEntity.AuthenticationEntity{}, err
	}

	password, err := auth.HashPassword(input.Password)
	if err != nil {
		logger.Error("hash password", "error", err)
		return &authEntity.AuthenticationEntity{}, err
	}

	user := &users.UserEntity{
		OriginalID: input.OriginalID,
		Email:      input.Email,
		Password:   password,
	}

	user, err = u.UserRepository.Create(ctx, user)
	if err != nil {
		logger.Error("Error create user", "input", user, "error", err)
		return &authEntity.AuthenticationEntity{}, err
	}

	claims := auth.NewClaims(user.ID)
	tokenString, err := claims.CreateTokenString()
	if err != nil {
		logger.Error("Error create jwt token", "claims", claims, "error", err)
		return &authEntity.AuthenticationEntity{}, err
	}

	return &authEntity.AuthenticationEntity{
		Token: tokenString,
		Exp:   claims.Exp,
		User:  *user,
	}, nil
}

func validateForSignup(r services.UserRepository, ctx context.Context, email string, originalID string) error {
	exist, err := r.ExistByEmail(ctx, email)
	if err != nil {
		return err
	}

	if exist {
		return shared.UserEmailExistsError.Error()
	}

	exist, err = r.ExistByOriginalID(ctx, originalID)
	if err != nil {
		return err
	}

	if exist {
		return shared.UserOriginalIDExistsError.Error()
	}

	return nil
}

func (u *userUseCase) Signin(ctx context.Context, input *input.SigninInput) (*authEntity.AuthenticationEntity, error) {
	var user *users.UserEntity
	var err error

	if utils.IsEmail(input.Username) {
		user, err = u.UserRepository.GetByEmail(ctx, input.Username)
	} else {
		user, err = u.UserRepository.GetByOriginalID(ctx, input.Username)
	}

	if err != nil {
		logger.Error("Error getting user", "error", err)
		return &authEntity.AuthenticationEntity{}, err
	}

	passwordMatch := auth.CheckPasswordHash(input.Password, user.Password)
	if !passwordMatch {
		return &authEntity.AuthenticationEntity{}, shared.UserPasswordNotMatchError.Error()
	}
	claims := auth.NewClaims(user.ID)
	tokenString, err := claims.CreateTokenString()
	if err != nil {
		logger.Error("Error create jwt token", "claims", claims, "error", err)
		return &authEntity.AuthenticationEntity{}, err
	}

	return &authEntity.AuthenticationEntity{
		Token: tokenString,
		Exp:   claims.Exp,
		User:  *user,
	}, nil
}
