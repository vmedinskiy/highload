package entity

import (
	"context"

	"github.com/go-faster/errors"
	"github.com/jmoiron/sqlx"

	"github.com/vmedinskiy/highload/internal/pkg/user"
	"github.com/vmedinskiy/highload/internal/pkg/user/repo"
)

type UserRepo interface {
	Create(ctx context.Context, user *user.User) (int64, error)
	GetByID(ctx context.Context, id int64) (*user.User, error)
	CheckPasswordByID(ctx context.Context, id int64, pass string) (*user.User, error)
	Search(ctx context.Context, firstName, secondName string) ([]*user.User, error)
}

type User struct {
	repo UserRepo
}

func NewUserEntity(db sqlx.ExtContext) *User {
	userRepo := repo.NewRepo(db, 8)
	return &User{
		repo: userRepo,
	}
}

func (u *User) Register(ctx context.Context, user *user.User) (int64, error) {
	id, err := u.repo.Create(ctx, user)
	if err != nil {
		return -1, errors.Wrap(err, "register new user")
	}
	return id, nil
}

func (u *User) GetByID(ctx context.Context, id int64) (*user.User, error) {
	resUser, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "get user by id")
	}
	return resUser, nil
}

func (u *User) Login(ctx context.Context, id int64, pass string) (*user.User, error) {
	resUser, err := u.repo.CheckPasswordByID(ctx, id, pass)
	if err != nil {
		return nil, errors.Wrap(err, "login user")
	}
	return resUser, nil
}

func (u *User) Search(ctx context.Context, firstName, lastName string) ([]*user.User, error) {
	res, err := u.repo.Search(ctx, firstName, lastName)
	if err != nil {
		return nil, errors.Wrap(err, "search users")
	}
	return res, nil
}
