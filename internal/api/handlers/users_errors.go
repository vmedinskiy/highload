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
		return (*api.GetUserApplicationJSONInternalServerError)(httperr.ApiError5xx(ctx, errResp))
	case http.StatusBadRequest:
		return (*api.GetUserApplicationJSONBadRequest)(httperr.ApiErrorGeneric(ctx, errResp))
	case http.StatusNotFound:
		return (*api.GetUserApplicationJSONNotFound)(httperr.ApiErrorGeneric(ctx, errResp))
	case http.StatusUnauthorized:
		return (*api.GetUserApplicationJSONUnauthorized)(httperr.ApiErrorGeneric(ctx, errResp))
	}
	return &api.GetUserApplicationJSONInternalServerError{}
}

func registerUserError(ctx context.Context, err error) api.RegisterUserRes {
	errResp := httperr.ToErrorResponder(err)
	switch errResp.Code {
	case http.StatusInternalServerError, http.StatusServiceUnavailable:
		return (*api.RegisterUserApplicationJSONInternalServerError)(httperr.ApiError5xx(ctx, errResp))
	case http.StatusBadRequest:
		return httperr.ApiErrorGeneric(ctx, errResp)
	}
	return &api.RegisterUserApplicationJSONInternalServerError{}
}

func loginUserError(ctx context.Context, err error) api.LoginUserRes {
	errResp := httperr.ToErrorResponder(err)
	switch errResp.Code {
	case http.StatusInternalServerError:
		return httperr.ApiError5xx(ctx, errResp)
	case http.StatusBadRequest:
		return (*api.LoginUserApplicationJSONBadRequest)(httperr.ApiErrorGeneric(ctx, errResp))
	case http.StatusNotFound:
		return (*api.LoginUserApplicationJSONNotFound)(httperr.ApiErrorGeneric(ctx, errResp))
	}
	return &api.Error5xxHeaders{}
}

func userSearchError(ctx context.Context, err error) api.UserSearchGetRes {
	errResp := httperr.ToErrorResponder(err)
	switch errResp.Code {
	case http.StatusBadRequest:
		return httperr.ApiErrorGeneric(ctx, errResp)
	case http.StatusInternalServerError, http.StatusServiceUnavailable:
		return (*api.UserSearchGetApplicationJSONInternalServerError)(httperr.ApiError5xx(ctx, errResp))
	}
	return &api.UserSearchGetApplicationJSONInternalServerError{}
}
