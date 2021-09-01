package service

import (
	"errors"
	"strings"
)

func (bot Bot) HandleRequest(update *Update) error {
	var response string
	var err error
	message := update.Message

	//Step 0: check if we are talking to a robot
	if message.From.IsBot {
		bot.log.Warn("This was a bot!")
		return errors.New("this is a bot")
	}

	//Step 1: trim first message word, the command.
	//		  This would direction the way the bot handles the request
	bot.log.Info("Sanitizing command")
	command := SanitizeCommand(message.Text)
	bot.log.Info("Command sanitized: " + command)

	//Step 2: switch command
	switch command {
	case "/start":
		response = PackInfo(message)
	case "/help":
		response = PackInfo(message)
	case "/generategame":
		response, err = GenerateUsername(strings.TrimPrefix(message.Text, command), "game")
		if err != nil {
			bot.log.Warn("Some game usernames could not have been created!")
			bot.log.WithError(err)
			response = PackError()
		}
	case "/generatedate":
		response, err = GenerateUsername(strings.TrimPrefix(message.Text, command), "date")
		if err != nil {
			bot.log.Warn("Some date usernames could not have been created!")
			bot.log.WithError(err)
			response = PackError()
		}
	case "/generateenterprise":
		response, err = GenerateUsername(strings.TrimPrefix(message.Text, command), "enterprise")
		if err != nil {
			bot.log.Warn("Some enterprise usernames could not have been created!")
			bot.log.WithError(err)
			response = PackError()
		}
	case "/generatemix":
		response, err = GenerateUsername(strings.TrimPrefix(message.Text, command), "mix")
		if err != nil {
			bot.log.Warn("Some usernames could not have been created!")
			bot.log.WithError(err)
			response = PackError()
		}
	case "/commands":
		response = PackCommands()
	default:
		response = PackDefault()
	}

	//Step 3: send usernames
	bot.log.Info("Sending response")
	err = bot.SendResponse(response, update.Message.Chat)
	if err != nil {
		bot.log.WithError(err).Error("There was an error sending the response back to the user!")
		return err
	}
	bot.log.Info("Response successfully sent!")

	return nil
}
