package handlers

import (
	"context"

	"github.com/go-faster/errors"

	api "github.com/vmedinskiy/highload/api/generated"
	"github.com/vmedinskiy/highload/internal/pkg/jwt"
	"github.com/vmedinskiy/highload/internal/pkg/usercontext"
)

// Compile-time check for Handler.
var _ api.SecurityHandler = (*SecHandler)(nil)

type SecHandler struct {
	jwtManager *jwt.Manager
	userEntity UserEntity
}

func NewSecHandler(userEntity UserEntity, jwtManager *jwt.Manager) *SecHandler {
	return &SecHandler{
		userEntity: userEntity,
		jwtManager: jwtManager,
	}
}

func (sc *SecHandler) HandleCookieAuth(ctx context.Context, _ string, t api.CookieAuth) (context.Context, error) {
	id, err := sc.jwtManager.Verify(t.GetAPIKey())
	if err != nil {
		return ctx, errors.Wrap(err, "hanlde cookie auth")
	}
	u, err := sc.userEntity.GetByID(ctx, id)
	if err != nil {
		return ctx, errors.Wrap(err, "get user during cookie handle auth")
	}
	usr := usercontext.User{
		ID:         u.ID,
		FirstName:  u.FirstName,
		SecondName: u.SecondName,
		Age:        u.Age.Int32,
		City:       u.City.String,
	}
	return usercontext.WrapUser(ctx, usr), nil
}
