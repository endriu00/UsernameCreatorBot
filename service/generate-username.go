package service

import (
	log "github.com/sirupsen/logrus"
	"strings"
)

func GenerateUsername(startingWord, requestType string) (string, error) {
	var usernames string

	//From here it starts the username creation methods invocation.
	if requestType == "mix" {
		log.Info("Generating usernames with vocal replacement. Function: ReplaceVocals")
		usernames = strings.Join([]string{usernames, ReplaceVocals(startingWord, "x", 2)}, "")
		log.Info("Function: ReplaceVocals exiting correctly. Usernames generated till now: " + usernames)
	}

	if requestType == "game" || requestType == "mix" {
		log.Info("Generating usernames with game names replacement. Function: ReplaceWithGameNames")
		usernameGames, err := ReplaceWithGameNames(startingWord, 4)
		if err != nil {
			log.WithError(err).Error("Failed to generate username with game names replacement")
			return usernames, err
		}
		usernames = strings.Join([]string{usernames, usernameGames}, "")
		log.Info("Function: ReplaceWithGameNames exiting correctly. Usernames generated till now: " + usernames)
	}

	if requestType == "date" || requestType == "mix" {
		log.Info("Generating usernames with date names replacement. Function: ReplaceWithDateNames")
		usernameDates, err := ReplaceWithDateNames(startingWord, 4)
		if err != nil {
			log.WithError(err).Error("Failed to generate username with date names replacement")
			return usernames, err
		}
		usernames = strings.Join([]string{usernames, usernameDates}, "")
		log.Info("Function: ReplaceWithDateNames exiting correctly. Usernames generated till now: " + usernames)
	}

	if requestType == "enterprise" || requestType == "mix" {
		log.Info("Generating usernames with enterprise names replacement. Function: ReplaceWithEnterpriseNames")
		usernameEnterps, err := ReplaceWithEnterpriseNames(startingWord, 4)
		if err != nil {
			log.WithError(err).Error("Failed to generate username with enterprise names replacement")
			return usernames, err
		}
		usernames = strings.Join([]string{usernames, usernameEnterps}, "")
		log.Info("Function: ReplaceWithEnterpriseNames exiting correctly. Usernames generated till now: " + usernames)
	}
	return usernames, nil
}
