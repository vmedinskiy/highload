package usercontext

import (
	"runtime"

	"github.com/go-faster/errors"
)

var (
	CtxIsEmptyErr     = errors.New("empty context")
	CtxWithoutUserErr = errors.Errorf("context does not contain %q", ContextUserKey)
)

func wrapFuncName(err error, skip int) error {
	if err == nil {
		return nil
	}

	if skip < 0 {
		skip = 0
	}

	pc, _, _, _ := runtime.Caller(skip + 2)
	return errors.Wrapf(err, "функция %q требует авторизации", runtime.FuncForPC(pc).Name())
}
