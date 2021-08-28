package service

import (
	"net/http"
	"net/url"
	"strconv"
)

func SendResponse(response string, chat Chat) error {
	telegramApiUrl := "https://api.telegram.org/bot1968866540:AAE8ZDEtUcEpwFqqLX9r5N3nrrQhHwwx4t4" + "/sendMessage"
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
