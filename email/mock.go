package email

import (
	"log"
)

func NewMock() Client {
	return &mockClient{}
}

type mockClient struct {
}

func (c *mockClient) Send(message Message) error {
	log.Println(message)

	return nil
}

func (c *mockClient) SendMultiple(recipients []Recipient, message Message) error {
	log.Println(message)

	return nil
}
