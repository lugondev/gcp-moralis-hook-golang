package mock

import (
	"log"
	"moralis-webhook/email"
)

type Config struct {
	Debug bool
}

func New(configuration Config) email.Client {
	return &mockClient{configuration: configuration}
}

type mockClient struct {
	configuration Config
}

func (c *mockClient) Send(message email.Message) error {
	if c.configuration.Debug {
		log.Println(message)
	}

	return nil
}
