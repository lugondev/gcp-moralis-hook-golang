package email

import (
	"github.com/mailjet/mailjet-apiv3-go"
	"log"
)

type MailjetConfig struct {
	Debug      bool
	PublicKey  string
	PrivateKey string
}

func NewMailjet(configuration MailjetConfig) Client {
	client := mailjet.NewMailjetClient(configuration.PublicKey, configuration.PrivateKey)

	return &mailjetClient{client: client, debug: configuration.Debug}
}

type mailjetClient struct {
	client *mailjet.Client
	debug  bool
}

func (c *mailjetClient) SendMultiple(recipients []Recipient, message Message) error {
	return nil
}

func (c *mailjetClient) Send(message Message) error {
	info := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Name:  message.Sender.Name,
				Email: message.Sender.Email,
			},
			To: &mailjet.RecipientsV31{
				{
					Name:  message.Recipient.Name,
					Email: message.Recipient.Email,
				},
			},
			Subject:  message.Subject,
			HTMLPart: message.HtmlContent,
			TextPart: message.TextContent,
		},
	}

	messages := &mailjet.MessagesV31{Info: info}

	response, err := c.client.SendMailV31(messages)

	if c.debug {
		log.Println(response)
		log.Println(err)
	}

	if err != nil {
		return err
	}

	return nil
}
