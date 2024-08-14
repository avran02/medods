package service

import (
	"fmt"
	"log/slog"

	"github.com/avran02/medods/config"
	"github.com/avran02/medods/internal/repository"
	"gopkg.in/gomail.v2"
)

type SMTPService interface {
	SendIPChangedNotification(userID, userIP string) error
}

type smtpService struct {
	mailer      *gomail.Dialer
	r           repository.Postgres
	IsAvailable bool
}

func (s *smtpService) SendIPChangedNotification(userID, userIP string) error {
	if !s.IsAvailable {
		slog.Warn("smtpService.SendIPChangedNotification: smtp is not available")
		return ErrSMTPUnavailable
	}
	m := gomail.NewMessage()
	email, _ := s.r.GetUserEmail(userID)
	m.SetHeader("From", s.mailer.Username, s.mailer.Username)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "IP Changed")
	m.SetBody("text/plain", fmt.Sprintf("IP с которого вы входили в последний раз не соответствует.\nПроизошла попытка авторизации с  IP: %s. Если это были не Вы, свяжитесь с нами.", userIP))

	return s.mailer.DialAndSend(m)
}

func NewSMTPService(config config.SMTPConfig, db repository.Postgres) SMTPService {
	mailer := gomail.NewDialer(
		config.Host,
		config.Port,
		config.User,
		config.Password,
	)
	if config.Host == "" || config.User == "" || config.Password == "" {
		slog.Warn("Cannot send emails. SMTP config is empty")
		return &smtpService{
			IsAvailable: false,
		}
	}
	return &smtpService{
		mailer:      mailer,
		r:           db,
		IsAvailable: true,
	}
}
