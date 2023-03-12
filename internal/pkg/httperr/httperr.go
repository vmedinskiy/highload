package httperr

import (
	"context"
	"net/http"

	"github.com/go-faster/errors"

	api "github.com/vmedinskiy/highload/api/generated"
)

type ErrorResponder struct {
	Msg        string
	Code       int
	RetryAfter int
	error
}

func ToErrorResponder(err error) *ErrorResponder {
	var targetError *ErrorResponder
	if errors.As(err, &targetError) {
		return targetError
	}
	//nolint:forcetypeassert
	return Wrap(err, http.StatusInternalServerError).(*ErrorResponder)
}

func ApiError5xx(_ context.Context, targetError *ErrorResponder) *api.Error5xxHeaders {
	result := &api.Error5xxHeaders{
		Response: api.Error5xx{
			Message: targetError.Msg,
			Code:    targetError.Code,
		},
	}
	if targetError.RetryAfter != 0 {
		result.Retryafter = api.NewOptInt(targetError.RetryAfter)
	}
	return result
}

func ApiErrorGeneric(_ context.Context, targetError *ErrorResponder) *api.ErrorGeneric {
	return &api.ErrorGeneric{
		Message: targetError.Msg,
		Code:    targetError.Code,
	}
}

func Wrap(err error, code int, retryAfter ...int) error {
	var rA int
	if len(retryAfter) > 0 {
		rA = retryAfter[0]
	}
	return &ErrorResponder{
		error:      err,
		Code:       code,
		Msg:        err.Error(),
		RetryAfter: rA,
	}
}

// bodyAllowedForStatus reports whether a given response status code
// permits a body. See RFC 7230, section 3.3.
func bodyAllowedForStatus(status int) bool {
	switch {
	case status >= 100 && status <= 199:
		return false
	case status == 204:
		return false
	case status == 304:
		return false
	}
	return true
}
