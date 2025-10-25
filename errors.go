package httperr

import "github.com/pkg/errors"

var (
	ErrInternalServer = errors.New("internal server error")
	ErrNotFound       = errors.New("not found")
	ErrRequestTimeout = errors.New("request timeout")
)
