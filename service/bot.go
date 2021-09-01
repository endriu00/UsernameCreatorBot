package service

import (
	"github.com/sirupsen/logrus"
)

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
