package sendgrid

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"moralis-webhook/email"
)

type Config struct {
	Debug  bool
	ApiKey string
}

func New(configuration Config) email.Client {
	client := sendgrid.NewSendClient(configuration.ApiKey)
	return &sendgridClient{client: client, debug: configuration.Debug}
}

type sendgridClient struct {
	client *sendgrid.Client
	debug  bool
}

func (c *sendgridClient) Send(message email.Message) error {
	from := mail.NewEmail(message.Sender.Name, message.Sender.Email)
	subject := message.Subject
	to := mail.NewEmail(message.Recipient.Name, message.Recipient.Email)

	m := mail.NewV3Mail()
	personalization := mail.NewPersonalization()
	personalization.AddTos(to)
	if message.CCs != nil && len(*message.CCs) > 0 {
		for _, cc := range *message.CCs {
			fmt.Println("CC: ", cc)
			personalization.AddCCs(mail.NewEmail(cc.Name, cc.Email))
		}
	}
	m.SetFrom(from)
	htmlContent := mail.NewContent("text/html", message.HtmlContent)
	plainContent := mail.NewContent("text/plain", message.TextContent)
	m.AddContent(plainContent, htmlContent)
	m.Subject = subject

	m.AddPersonalizations(personalization)
	response, err := c.client.Send(m)

	if c.debug {
		log.Println(response)
		log.Println(err)
	}

	if err != nil {
		return err
	}

	return nil
}
