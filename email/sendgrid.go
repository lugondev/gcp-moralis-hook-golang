package email

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"moralis-webhook/notifier"
	"sync"
)

type SendgridConfig struct {
	Debug  bool
	ApiKey string
}

func NewSendgrid(config SendgridConfig, notifier *notifier.Notifier) Client {
	client := sendgrid.NewSendClient(config.ApiKey)
	return &sendgridClient{client: client, debug: config.Debug, notifier: notifier}
}

type sendgridClient struct {
	client   *sendgrid.Client
	debug    bool
	notifier *notifier.Notifier
}

func (c *sendgridClient) notify(message notifier.Message) {
	if c.notifier != nil {
		err := (*c.notifier).Notify(message)
		if err != nil {
			log.Println("failed to notify: ", err)
		}
	}
}

func (c *sendgridClient) send(message Message) error {
	from := mail.NewEmail(message.Sender.Name, message.Sender.Email)
	subject := message.Subject
	to := mail.NewEmail(message.Recipient.Name, message.Recipient.Email)

	m := mail.NewV3Mail()
	personalization := mail.NewPersonalization()
	personalization.AddTos(to)
	if message.CCs != nil && len(*message.CCs) > 0 {
		for _, cc := range *message.CCs {
			log.Println("CC: ", cc)
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
	if err != nil {
		log.Printf("failed to send email to: %s. err: %v\n", message.Recipient.Email, err)
		c.notify(notifier.Message{
			Content: fmt.Sprintf("failed to send email to: %s. err: %v\n", message.Recipient.Email, err),
		})
		return err
	}

	if c.debug {
		log.Println(response)
	}

	return nil
}

func (c *sendgridClient) Send(message Message) error {
	if message.Recipients != nil && len(*message.Recipients) > 0 {
		go c.SendMultiple(*message.Recipients, message)
		return nil
	}
	return c.send(message)
}

func (c *sendgridClient) SendMultiple(recipients []Recipient, message Message) int {
	var wg sync.WaitGroup
	wg.Add(len(recipients))
	go func() {
		for _, recipient := range recipients {
			message.Recipient = recipient
			_ = c.send(message)
			wg.Done()
		}
	}()
	wg.Wait()
	return len(recipients)
}
