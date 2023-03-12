package usercontext

import (
	"context"

	"github.com/go-faster/errors"
)

type User struct {
	ID         int64
	FirstName  string
	SecondName string
	Age        int32
	City       string
}

type contextIdentityKey int

const (
	_ contextIdentityKey = iota
	ContextUserKey
)

// WrapUser возвращает context.Context, с хранимым в нем db.User.
func WrapUser(ctx context.Context, user User) context.Context {
	return context.WithValue(ctx, ContextUserKey, user)
}

func (key contextIdentityKey) String() string {
	switch key {
	case ContextUserKey:
		return "context user key"
	default:
		return "unknown key"
	}
}

// GetUser called after authHandler, error if it's not
func GetUser(ctx context.Context) (u User, err error) {
	defer func() { err = wrapFuncName(err, 1) }()

	if ctx == nil {
		return User{}, ErrCtxIsEmpty
	}

	cuk := ctx.Value(ContextUserKey)
	if cuk == nil {
		return User{}, ErrCtxWithoutUser
	}

	switch cuk := cuk.(type) {
	case User:
		return cuk, nil
	case error:
		return User{}, cuk
	}

	return User{}, errors.Errorf("unexpected value for user %v (%[1]T), ожидается %T", cuk, User{})
}
