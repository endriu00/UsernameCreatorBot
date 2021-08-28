package service

import (
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func SendResponse(response string, chat Chat) error {
	telegramApiUrl := "https://api.telegram.org/bot" + os.Getenv("TOKEN") + "/sendMessage"
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
