package handler

import (
	"hiyoko-fiber/internal/application/usecase"
	"hiyoko-fiber/internal/presentation/http/app/input"
	"hiyoko-fiber/internal/presentation/http/app/oapi"
	"hiyoko-fiber/internal/shared"
	"hiyoko-fiber/pkg/logging/file"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	Signup(c *fiber.Ctx) error
	Signin(c *fiber.Ctx) error
}

type authHandler struct {
	UserUseCase usecase.UserUseCase
}

func NewAuthHandler(u usecase.UserUseCase) AuthHandler {
	return &authHandler{u}
}

func (h *authHandler) Signup(c *fiber.Ctx) error {
	ctx := c.Context()
	userInput := new(input.SignupInput)
	if err := c.BodyParser(userInput); err != nil {
		logger.Error("Error BodyParser signup input", "userInput", userInput, "error", err)
		return shared.ResponseBadRequest(c, shared.NoneCode)
	}

	if err := userInput.Validate(); err != nil {
		return shared.ResponseBadRequest(c, shared.NoneCode)
	}

	authEntity, err := h.UserUseCase.Signup(ctx, userInput)
	if err != nil {
		return shared.ResponseBadRequest(c, shared.NoneCode)
	}

	return shared.ResponseCreate(c, oapi.SignupResponse{
		Authorisation: oapi.Authorisation{
			Token: authEntity.Token,
			Exp:   authEntity.Exp,
		},
		User: oapi.UserSchema{
			ID:         authEntity.User.ID,
			OriginalID: &authEntity.User.OriginalID,
			Email:      authEntity.User.Email,
		},
	})
}

func (h *authHandler) Signin(c *fiber.Ctx) error {
	ctx := c.Context()
	userInput := new(input.SigninInput)
	if err := c.BodyParser(userInput); err != nil {
		logger.Error("Error BodyParser signin input", "userInput", userInput, "error", err)
		return shared.ResponseBadRequest(c, shared.NoneCode)
	}

	if err := userInput.Validate(); err != nil {
		return shared.ResponseBadRequest(c, shared.NoneCode)
	}

	authEntity, err := h.UserUseCase.Signin(ctx, userInput)
	if err != nil {
		return shared.ResponseUnauthorized(c)
	}

	return shared.ResponseCreate(c, oapi.SigninResponse{
		Authorisation: oapi.Authorisation{
			Token: authEntity.Token,
			Exp:   authEntity.Exp,
		},
		User: oapi.UserSchema{
			ID:         authEntity.User.ID,
			OriginalID: &authEntity.User.OriginalID,
			Email:      authEntity.User.Email,
		},
	})
}
