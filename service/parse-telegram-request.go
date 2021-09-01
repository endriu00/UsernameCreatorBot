package service

import (
	"encoding/json"
	"net/http"
)

func (bot *Bot) ParseTelegramRequest(r *http.Request) (*Update, error) {
	var update Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		bot.log.Printf("Incoming update cannot be decoded! %s", err.Error())
		bot.log.Print(update)
	}
	return &update, nil
}
