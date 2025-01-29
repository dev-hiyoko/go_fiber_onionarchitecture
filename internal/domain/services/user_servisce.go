package services

import (
	"context"

	"hiyoko-fiber/internal/domain/entities/users"
	"hiyoko-fiber/internal/pkg/ent/util"
)

type UserRepository interface {
	Get(ctx context.Context, id *util.ULID) (*users.UserEntity, error)
	GetByEmail(ctx context.Context, email string) (*users.UserEntity, error)
	GetByOriginalID(ctx context.Context, originalID string) (*users.UserEntity, error)
	ExistByEmail(ctx context.Context, email string) (bool, error)
	ExistByOriginalID(ctx context.Context, originalID string) (bool, error)
	Create(ctx context.Context, user *users.UserEntity) (*users.UserEntity, error)
	Update(ctx context.Context, user *users.UserEntity) (*users.UserEntity, error)
}
