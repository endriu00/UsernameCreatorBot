package main

import (
	"github.com/endriu00/UsernameCreatorBot/service"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", http.HandlerFunc(service.HandleTelegramWebhook))
}
