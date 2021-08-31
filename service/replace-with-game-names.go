package service

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

func ReplaceWithGameNames(startingWord string, repNum int) (string, error) {
	var usernames string
	var hits int = 0
	file, err := os.Open("names/game.txt")
	if err != nil {
		return usernames, err
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if hits == repNum {
			break
		}
		idx := 0
		currentWord := ""
		for _, token := range startingWord {

			log.Info(currentWord)

			idx++
			letter := string(token)
			currentWord = strings.Join([]string{currentWord, letter}, "")
			if idx < 2 {
				continue
			}
			if strings.HasPrefix(scanner.Text(), letter) {
				gameWord := strings.Join([]string{strings.TrimSuffix(currentWord, letter), scanner.Text()}, "")
				gameWord = strings.TrimSpace(gameWord)
				usernames = strings.Join([]string{usernames, gameWord}, "\n")
				hits++
			}
		}
	}
	return usernames, nil
}
