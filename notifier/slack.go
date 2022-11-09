package notifier

import (
	"github.com/labstack/gommon/log"
	"github.com/slack-go/slack"
)

type SlackConfig struct {
	WebhookUrl string
}

func NewSlack(configuration SlackConfig) *SlackClient {
	return &SlackClient{configuration}
}

type SlackClient struct {
	config SlackConfig
}

func (c *SlackClient) Notify(message Message) error {
	webhook := slack.WebhookMessage{
		Text: message.Content,
	}

	if err := slack.PostWebhook(c.config.WebhookUrl, &webhook); err != nil {
		log.Error(err)
		return err
	}
	return nil
}
