package input

import (
	"github.com/go-playground/validator/v10"
)

type SignupInput struct {
	OriginalID string `json:"originalID" validate:"omitempty,min=3,max=36"`
	Email      string `json:"email" validate:"required,email,max=255"`
	Password   string `json:"password" validate:"required,min=8,max=72"`
}

func (i SignupInput) Validate() error {
	return validator.New().Struct(i)
}

type SigninInput struct {
	Username string `json:"username" validate:"required,max=255"`
	Password string `json:"password" validate:"required,max=72"`
}

func (i SigninInput) Validate() error {
	return validator.New().Struct(i)
}
