package service

import (
	log "github.com/sirupsen/logrus"
	"strings"
)

func GenerateUsername(startingWord string) (string, error) {
	var usernames string

	//From here it starts the username creation methods invocation.
	log.Info("Generating usernames with vocal replacement. Function: ReplaceVocals")
	usernames = strings.Join([]string{usernames, ReplaceVocals(startingWord, "x", 2)}, "")
	log.Info("Function: ReplaceVocals exiting correctly. Usernames generated till now: " + usernames)

	log.Info("Generating usernames with game names replacement. Function: ReplaceWithGameNames")
	usernameGames, err := ReplaceWithGameNames(startingWord, 4)
	if err != nil {
		log.WithError(err).Error("Failed to generate username with game names replacement")
		return usernames, err
	}
	usernames = strings.Join([]string{usernames, usernameGames}, "")
	log.Info("Function: ReplaceWithGameNames exiting correctly. Usernames generated till now: " + usernames)

	return usernames, nil
}
