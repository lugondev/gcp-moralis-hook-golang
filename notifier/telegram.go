package notifier

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type TelegramConfig struct {
	ApiToken string
	Debug    bool
	ChatId   int64
}

func NewTelegram(configuration TelegramConfig) *TelegramClient {
	bot, err := tgbotapi.NewBotAPI(configuration.ApiToken)
	if err != nil {
		log.Fatal("Unable to initialize telegram bot:", err)
	}
	bot.Debug = configuration.Debug
	return &TelegramClient{configuration, bot}
}

type TelegramClient struct {
	config TelegramConfig
	bot    *tgbotapi.BotAPI
}

func (c *TelegramClient) Notify(message Message) error {
	msg := tgbotapi.NewMessage(c.config.ChatId, message.Content)
	msg.ParseMode = tgbotapi.ModeHTML

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Println("Unable to send message telegram:", err)
		return err
	}

	return nil
}
