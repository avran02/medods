package controller

import "github.com/avran02/medods/internal/service"

type Controller interface {
	AuthenticationController
}

type controller struct {
	AuthenticationController
}

func New(s service.Service) Controller {
	return &controller{
		AuthenticationController: NewAuthenticationController(s),
	}
}
