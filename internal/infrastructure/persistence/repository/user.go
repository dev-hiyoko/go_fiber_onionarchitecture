package repository

import (
	"context"

	"hiyoko-fiber/internal/domain/entities/users"
	"hiyoko-fiber/internal/domain/services"
	"hiyoko-fiber/internal/infrastructure/database"
	"hiyoko-fiber/internal/pkg/ent/models"
	"hiyoko-fiber/internal/pkg/ent/user"
	"hiyoko-fiber/internal/pkg/ent/util"
)

type userRepository struct {
	conn *database.MysqlEntClient
}

func NewUserRepository(conn *database.MysqlEntClient) services.UserRepository {
	return &userRepository{conn}
}

func (r *userRepository) Get(ctx context.Context, id *util.ULID) (*users.UserEntity, error) {
	u, err := r.conn.User.Get(ctx, *id)
	if err != nil {
		return nil, err
	}
	return models.UsersModel{User: *u}.ToEntity(), nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*users.UserEntity, error) {
	u, err := r.conn.User.
		Query().
		Where(user.EmailEqualFold(email)).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return models.UsersModel{User: *u}.ToEntity(), nil
}

func (r *userRepository) GetByOriginalID(ctx context.Context, originalID string) (*users.UserEntity, error) {
	u, err := r.conn.User.
		Query().
		Where(user.OriginalIDEqualFold(originalID)).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return models.UsersModel{User: *u}.ToEntity(), nil
}

func (r *userRepository) ExistByEmail(ctx context.Context, email string) (bool, error) {
	exist, err := r.conn.User.
		Query().
		Where(user.EmailEqualFold(email)).
		Exist(ctx)
	if err != nil {
		return exist, err
	}
	return exist, nil
}

func (r *userRepository) ExistByOriginalID(ctx context.Context, originalID string) (bool, error) {
	exist, err := r.conn.User.
		Query().
		Where(user.OriginalIDEqualFold(originalID)).
		Exist(ctx)
	if err != nil {
		return exist, err
	}
	return exist, nil
}

func (r *userRepository) Create(ctx context.Context, u *users.UserEntity) (*users.UserEntity, error) {
	userClient := r.conn.User.Create().
		SetEmail(u.Email).
		SetPassword(u.Password)

	if u.OriginalID != "" {
		userClient.SetOriginalID(u.OriginalID)
	}

	entUser, err := userClient.Save(ctx)

	if err != nil {
		return nil, err
	}
	return models.UsersModel{User: *entUser}.ToEntity(), nil
}

func (r *userRepository) Update(ctx context.Context, u *users.UserEntity) (*users.UserEntity, error) {
	entUser, err := r.conn.User.Get(ctx, util.ULID(u.ID))
	if err != nil {
		return nil, err
	}

	entUser, err = entUser.Update().
		SetStatus(u.Status).
		SetOriginalID(u.OriginalID).
		SetEmail(u.Email).
		SetPassword(u.Password).
		SetUpdatedAt(u.UpdatedAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return models.UsersModel{User: *entUser}.ToEntity(), nil
}
