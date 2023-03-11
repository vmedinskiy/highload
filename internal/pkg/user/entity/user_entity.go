package entity

import (
	"context"

	"github.com/go-faster/errors"
	"github.com/jmoiron/sqlx"

	"github.com/vmedinskiy/highload/internal/pkg/user"
	"github.com/vmedinskiy/highload/internal/pkg/user/repo"
)

type UserRepo interface {
	Create(ctx context.Context, user user.User) (int64, error)
	GetByID(ctx context.Context, id int64) (*user.User, error)
	CheckPasswordByID(ctx context.Context, id int64, pass string) (*user.User, error)
}

type UserEntity struct {
	repo UserRepo
}

func NewUserEntity(db sqlx.ExtContext) *UserEntity {
	userRepo := repo.NewRepo(db, 8)
	return &UserEntity{
		repo: userRepo,
	}
}

func (u *UserEntity) Register(ctx context.Context, user user.User) (int64, error) {
	id, err := u.repo.Create(ctx, user)
	return id, errors.Wrap(err, "register new user")
}

func (u *UserEntity) GetByID(ctx context.Context, id int64) (*user.User, error) {
	resUser, err := u.repo.GetByID(ctx, id)
	return resUser, errors.Wrap(err, "get user by id")
}

func (u *UserEntity) Login(ctx context.Context, id int64, pass string) (*user.User, error) {
	resUser, err := u.repo.CheckPasswordByID(ctx, id, pass)
	return resUser, errors.Wrap(err, "check password by id")
}
