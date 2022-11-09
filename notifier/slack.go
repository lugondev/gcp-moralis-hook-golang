package notifier

import (
	"github.com/labstack/gommon/log"
	"github.com/slack-go/slack"
)

type SlackConfig struct {
	WebhookUrl string
}

func NewSlack(configuration SlackConfig) *SlackClient {
	return &SlackClient{webhookUrl: configuration.WebhookUrl}
}

type SlackClient struct {
	webhookUrl string
}

func (c *SlackClient) Notify(message Message) {
	webhook := slack.WebhookMessage{
		Text: message.Content,
	}

	if err := slack.PostWebhook(c.webhookUrl, &webhook); err != nil {
		log.Error(err)
	}
}
