package service

import (
	"github.com/avran02/medods/config"
	"github.com/avran02/medods/internal/pkg/jwt"
	"github.com/avran02/medods/internal/repository"
)

type Service interface {
	AuthenticationService
	SMTPService
}

type service struct {
	AuthenticationService
	SMTPService
}

func New(r repository.Repository, j jwt.JwtGenerator, smtpConf config.SMTPConfig) Service {
	return &service{
		AuthenticationService: NewAuthenticationService(r, j),
		SMTPService:           NewSMTPService(smtpConf, r),
	}
}
