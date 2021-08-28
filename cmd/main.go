package main

import (
	"github.com/endriu00/UsernameCreatorBot/service"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(service.HandleTelegramWebhook))
}
