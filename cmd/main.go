package main

import (
	"github.com/endriu00/UsernameCreatorBot/service"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {

	//Start logger
	log := logrus.NewEntry(logrus.StandardLogger())

	//Load configuration
	config, err := loadConfiguration()
	if err != nil {
		log.Error("Failed loading configuration.")
		//maybe return something, create another function that main can call to handle errors
	}

	//Create bot
	bot := service.New(&service.Config{
		Logger:           config.Logger,
		TelegramApiUrl:   config.TelegramApiUrl,
		TelegramBotToken: config.TelegramBotToken,
	})

	http.ListenAndServe(":"+os.Getenv(`PORT`), http.HandlerFunc(bot.HandleTelegramWebhook))
}
