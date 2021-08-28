package service

import (
	"net/http"
)

func HandleTelegramWebhook(r *http.Request) {

	update, err := ParseTelegramRequest(r)
	if err != nil {
		//bot.log.WithError(err).Error("Cannot parse telegram request.")
		return
	}
	err = HandleRequest(update)
	if err != nil {
		//bot.log.WithError(err).Error("Handler returned an error.")
		return
	}
}
