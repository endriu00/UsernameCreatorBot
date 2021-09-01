package service

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

type BotInt interface {
	HandleTelegramWebhook(w http.ResponseWriter, r *http.Request)
}

type Bot struct {
	log              *logrus.Entry
	telegramApiUrl   string
	telegramBotToken string
}

type Config struct {
	Logger           *logrus.Entry
	TelegramApiUrl   string
	TelegramBotToken string
}

func New(conf *Config) Bot {
	return Bot{
		log:              conf.Logger,
		telegramApiUrl:   conf.TelegramApiUrl,
		telegramBotToken: conf.TelegramBotToken,
	}
}
