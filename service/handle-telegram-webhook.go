package service

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func HandleTelegramWebhook(w http.ResponseWriter, r *http.Request) {

	update, err := ParseTelegramRequest(r)
	if err != nil {
		log.WithError(err).Error("Cannot parse telegram request.")
		return
	}
	err = HandleRequest(update)
	if err != nil {
		log.WithError(err).Error("Handler returned an error.")
		return
	}
}
