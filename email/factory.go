package email

import (
	"errors"
	"github.com/spf13/viper"
	"log"
	"moralis-webhook/notifier"
)

type Adapter string

const (
	Mock     Adapter = "mock"
	Sendgrid Adapter = "sendgrid"
	Mailjet  Adapter = "mailjet"
)

type Config struct {
	Adapter  Adapter
	Notifier bool
	Sendgrid SendgridConfig
	Mailjet  MailjetConfig
}

func NewEmailClient(notifier *notifier.Notifier) (Client, error) {
	var config Config
	err := viper.UnmarshalKey("email", &config)
	if err != nil {
		return nil, err
	}

	switch config.Adapter {
	case Mock:
		log.Println("Initialize Mock Email Client")
		return NewMock(), nil
	case Sendgrid:
		log.Println("Initialize Sendgrid Email Client")
		if config.Notifier {
			return NewSendgrid(config.Sendgrid, notifier), nil
		} else {
			return NewSendgrid(config.Sendgrid, nil), nil
		}
	case Mailjet:
		log.Println("Initialize Mailjet Email Client")
		return NewMailjet(config.Mailjet), nil
	default:
		return nil, errors.New("unable to find adapter")
	}
}
