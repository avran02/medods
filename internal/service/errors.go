package service

import (
	"errors"
)

var (
	ErrWrongTokensPair = errors.New("access token and refresh token are not pair")
	ErrUnknownToken    = errors.New("unknown token")
)
