package factory

import (
	"errors"
	"log"
	"moralis-webhook/email"
	"moralis-webhook/email/mailjet"
	"moralis-webhook/email/mock"
	"moralis-webhook/email/sendgrid"
)

type Config struct {
	Adapter  string
	Mock     mock.Config
	Mailjet  mailjet.Config
	Sendgrid sendgrid.Config
}

func NewEmailClient(config Config) (email.Client, error) {
	switch config.Adapter {
	case "mock":
		log.Println("Initialize Mock Email Client")
		return mock.New(config.Mock), nil
	case "sendgrid":
		log.Println("Initialize Sendgrid Email Client")
		return sendgrid.New(config.Sendgrid), nil
	case "mailjet":
		log.Println("Initialize Mailjet Email Client")
		return mailjet.New(config.Mailjet), nil
	default:
		return nil, errors.New("unable to find adapter")
	}
}
