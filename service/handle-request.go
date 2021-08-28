package service

import (
	"errors"
	"strings"
)

//TODO: add logging
func HandleRequest(update *Update) error {
	var response string
	var err error
	message := update.Message

	//Step 0: check if we are talking to a robot
	if message.From.IsBot {
		return errors.New("this is a bot")
	}

	//Step 1: trim first message word, the command.
	//		  This would direction the way the bot handles the request
	command := SanitizeCommand(message.Text)

	//Step 2: switch command
	switch command {
	case "/start":
		response = PackInfo(message)
	case "/help":
		response = PackInfo(message)

	case "/generate":
		response = GenerateUsername(strings.TrimPrefix(message.Text, command))
	case "/commands":
		response = PackCommands()
	default:
		response = PackCommands()
	}

	//Step 3: send usernames
	err = SendResponse(response, update.Message.Chat)
	if err != nil {
		return err
	}

	return nil
}
