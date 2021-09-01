package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

type BotConfiguration struct {
	Logger *logrus.Entry

	TelegramApiUrl   string
	TelegramBotToken string
}

func loadConfiguration() (BotConfiguration, error) {
	return BotConfiguration{
		TelegramApiUrl:   "https://api.telegram.org/bot",
		TelegramBotToken: os.Getenv(`TOKEN`),

		Logger: nil,
	}, nil
}
