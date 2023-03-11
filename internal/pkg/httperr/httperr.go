package httperr

import (
	"encoding/json"
	"errors"
	"net/http"
)

const (

	// HeaderAccept http header Accept.
	HeaderAccept      = "Accept"
	HeaderContentType = "Content-Type"
	// ContentTypeJSON content type json.
	ContentTypeJSON = "application/json"
)

var unexpectedNilError = responder{
	Msg:  "unexpected nil error",
	Code: http.StatusInternalServerError,
}

type responder struct {
	Msg  string `json:"message"`
	Code int    `json:"code"`
	error
}

func (v responder) writeResponse(w http.ResponseWriter) {
	w.Header().Add(HeaderAccept, ContentTypeJSON)
	w.Header().Add(HeaderContentType, ContentTypeJSON)
	w.WriteHeader(v.Code)
	if bodyAllowedForStatus(v.Code) {
		payload, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}

		if _, err := w.Write(payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func Response(w http.ResponseWriter, err error, code int) {
	if err == nil {
		unexpectedNilError.writeResponse(w)
		return
	}

	var targetError *responder
	if errors.As(err, &targetError) {
		targetError.writeResponse(w)
		return
	}

	resp := responder{
		Msg:  err.Error(),
		Code: code,
	}
	resp.writeResponse(w)
}

func Wrap(err error, code int) error {
	return &responder{
		error: err,
		Code:  code,
		Msg:   err.Error(),
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
