package service

import (
	"errors"
)

var (
	ErrWrongTokensPair = errors.New("access token and refresh token are not pair")
	ErrUnknownToken    = errors.New("unknown token")
	// ErrIPChanged       = errors.New("ip changed")
	ErrSMTPUnavailable = errors.New("smtp is not available")
)
