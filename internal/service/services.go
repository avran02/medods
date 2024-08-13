package service

import (
	"github.com/avran02/medods/internal/pkg/jwt"
	"github.com/avran02/medods/internal/repository"
)

type Service interface {
	AuthenticationService
}

type service struct {
	AuthenticationService
}

func New(r repository.Repository, j jwt.JwtGenerator) Service {
	return &service{
		AuthenticationService: NewAuthenticationService(r, j),
	}
}
