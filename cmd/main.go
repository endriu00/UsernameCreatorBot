package main

import (
	"github.com/endriu00/UsernameCreatorBot/service"
	"net/http"
	"os"
)

func main() {
	http.ListenAndServe(":"+os.Getenv(`PORT`), http.HandlerFunc(service.HandleTelegramWebhook))
}
