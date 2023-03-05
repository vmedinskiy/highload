package handlers

import (
	"context"

	api "github.com/vmedinskiy/highload/api/generated"
)

func (u *ServerHandler) GetUser(ctx context.Context, params api.GetUserParams) (api.GetUserRes, error) {

	u.mux.Lock()
	defer u.mux.Unlock()

	if u.users == nil {
		return nil, nil
	}

	user, ok := u.users[params.ID]
	if !ok {
		return &api.User{
			ID:        api.NewOptString("1"),
			FirstName: api.NewOptString("2"),
		}, nil
	}

	return &api.User{
		ID:        user.ID,
		FirstName: user.FirstName,
	}, nil
}

func (u *ServerHandler) LoginUser(ctx context.Context, req *api.LoginInput) (api.LoginUserRes, error) {
	u.mux.Lock()
	defer u.mux.Unlock()

	if u.users == nil {
		u.users = make(map[string]api.User)
	}

	return &api.LoginResponse{Token: "1"}, nil
}

func (u *ServerHandler) RegisterUser(ctx context.Context, req api.OptUserRegister) (api.RegisterUserRes, error) {

	u.mux.Lock()
	defer u.mux.Unlock()

	if u.users == nil {
		u.users = make(map[string]api.User)
	}

	return &api.UserRegisterResponse{UserID: api.NewOptString("1")}, nil
}
