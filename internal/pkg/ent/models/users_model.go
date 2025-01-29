package models

import (
	"hiyoko-fiber/internal/domain/entities/users"
	"hiyoko-fiber/internal/pkg/ent"
)

type UsersModel struct {
	ent.User
}

func (u UsersModel) ToEntity() *users.UserEntity {
	return &users.UserEntity{
		ID:         u.ID.String(),
		Status:     u.Status,
		OriginalID: u.OriginalID,
		Email:      u.Email,
		Password:   u.Password,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}
