package email

import (
	"fmt"
	"net/smtp"
)

type Email interface {
	SendEmail(to []string, message []byte) error
}

type ServiceConfig struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

type email struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

func NewEmailService(cfg ServiceConfig) Email {
	return &email{
		Email:    cfg.Email,
		Password: cfg.Password,
		Host:     cfg.Host,
		Port:     cfg.Port,
	}
}

func (e *email) SendEmail(to []string, message []byte) error {
	auth := smtp.PlainAuth("", e.Email, e.Password, e.Host)
	err := smtp.SendMail(e.address(), auth, e.Email, to, message)

	if err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}

	return nil
}

func (e *email) address() string {
	return e.Host + ":" + e.Port
}
