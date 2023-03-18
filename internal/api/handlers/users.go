package handlers

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	api "github.com/vmedinskiy/highload/api/generated"
	"github.com/vmedinskiy/highload/internal/pkg/user"
)

func (s *ServerHandler) GetUser(ctx context.Context, params api.GetUserParams) (api.GetUserRes, error) {
	u, err := s.userEntity.GetByID(ctx, params.ID)
	if err != nil {
		return getUserError(ctx, err), nil
	}
	return &api.User{
		ID:         u.ID,
		FirstName:  u.FirstName,
		SecondName: u.SecondName,
		Age:        u.Age.Int32,
		Biography:  u.Biography.String,
		City:       u.City.String,
	}, nil
}

func (s *ServerHandler) LoginUser(ctx context.Context, req *api.LoginInput) (api.LoginUserRes, error) {
	u, err := s.userEntity.Login(ctx, req.ID, req.Password)
	if err != nil {
		return loginUserError(ctx, err), nil
	}
	t, err := s.jwtManager.Generate(u)
	if err != nil {
		return loginUserError(ctx, err), nil
	}
	cookie := http.Cookie{
		Name:     "xsession",
		Value:    t,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	return &api.LoginResponseHeaders{
		Response: api.LoginResponse{
			Token: t,
		},
		SetCookie: cookie.String(),
	}, nil
}

func (s *ServerHandler) RegisterUser(ctx context.Context, req *api.UserRegister) (api.RegisterUserRes, error) {
	id, err := s.userEntity.Register(ctx, &user.User{
		FirstName:  req.FirstName,
		SecondName: req.SecondName,
		Age:        sql.NullInt32{Int32: req.Age.Value, Valid: req.Age.Set},
		Biography:  sql.NullString{String: req.Biography.Value, Valid: req.Biography.Set},
		City:       sql.NullString{String: req.City.Value, Valid: req.City.Set},
		Password:   req.Password,
	})

	if err != nil {
		return registerUserError(ctx, err), nil
	}

	return &api.UserRegisterResponse{UserID: id}, nil
}
