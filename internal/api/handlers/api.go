package handlers

import (
	"context"

	api "github.com/vmedinskiy/highload/api/generated"
	"github.com/vmedinskiy/highload/internal/pkg/jwt"
	"github.com/vmedinskiy/highload/internal/pkg/user"
)

// Compile-time check for Handler.
var _ api.Handler = (*ServerHandler)(nil)

type UserEntity interface {
	Register(ctx context.Context, user *user.User) (int64, error)
	GetByID(ctx context.Context, id int64) (*user.User, error)
	Login(ctx context.Context, id int64, pass string) (*user.User, error)
}

type ServerHandler struct {
	userEntity UserEntity
	jwtManager *jwt.Manager
}

func NewServerHandler(userEntity UserEntity, jwtManager *jwt.Manager) *ServerHandler {
	return &ServerHandler{
		userEntity: userEntity,
		jwtManager: jwtManager,
	}
}
