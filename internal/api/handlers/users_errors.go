//nolint:forcetypeassert
package handlers

import (
	"context"
	"net/http"

	api "github.com/vmedinskiy/highload/api/generated"
	"github.com/vmedinskiy/highload/internal/pkg/httperr"
)

func getUserError(ctx context.Context, err error) api.GetUserRes {
	errResp := httperr.ToErrorResponder(err)
	switch errResp.Code {
	case http.StatusInternalServerError, http.StatusServiceUnavailable:
		return (*api.GetUserInternalServerError)(httperr.ApiError5xx(ctx, errResp))
	case http.StatusBadRequest:
		return (*api.GetUserBadRequest)(httperr.ApiErrorGeneric(ctx, errResp))
	case http.StatusNotFound:
		return (*api.GetUserNotFound)(httperr.ApiErrorGeneric(ctx, errResp))
	case http.StatusUnauthorized:
		return (*api.GetUserUnauthorized)(httperr.ApiErrorGeneric(ctx, errResp))
	}
	return &api.GetUserInternalServerError{}
}

func registerUserError(ctx context.Context, err error) api.RegisterUserRes {
	errResp := httperr.ToErrorResponder(err)
	switch errResp.Code {
	case http.StatusInternalServerError, http.StatusServiceUnavailable:
		return (*api.RegisterUserInternalServerError)(httperr.ApiError5xx(ctx, errResp))
	case http.StatusBadRequest:
		return httperr.ApiErrorGeneric(ctx, errResp)
	}
	return &api.RegisterUserInternalServerError{}
}

func loginUserError(ctx context.Context, err error) api.LoginUserRes {
	errResp := httperr.ToErrorResponder(err)
	switch errResp.Code {
	case http.StatusInternalServerError:
		return httperr.ApiError5xx(ctx, errResp)
	case http.StatusBadRequest:
		return (*api.LoginUserBadRequest)(httperr.ApiErrorGeneric(ctx, errResp))
	case http.StatusNotFound:
		return (*api.LoginUserNotFound)(httperr.ApiErrorGeneric(ctx, errResp))
	}
	return &api.Error5xxHeaders{}
}

func userSearchError(ctx context.Context, err error) api.UserSearchGetRes {
	errResp := httperr.ToErrorResponder(err)
	switch errResp.Code {
	case http.StatusBadRequest:
		return httperr.ApiErrorGeneric(ctx, errResp)
	case http.StatusInternalServerError, http.StatusServiceUnavailable:
		return (*api.UserSearchGetInternalServerError)(httperr.ApiError5xx(ctx, errResp))
	}
	return &api.UserSearchGetInternalServerError{}
}
