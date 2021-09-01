package service

import (
	"net/http"
	"net/url"
	"strconv"
)

func (bot *Bot) SendResponse(response string, chat Chat) error {
	telegramApiUrl := bot.telegramApiUrl + bot.telegramBotToken + "/sendMessage"
	_, err := http.PostForm(
		telegramApiUrl,
		url.Values{
			"chat_id": {strconv.Itoa(chat.Id)},
			"text":    {response},
		})
	if err != nil {
		return err
	}
	return nil
}
