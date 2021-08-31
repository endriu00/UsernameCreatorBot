package service

import (
	"bufio"
	"os"
	"strings"
)

func ReplaceWithGameNames(startingWord string, repNum int) (string, error) {
	var usernames string
	file, err := os.Open("names/game.txt")
	if err != nil {
		return usernames, err
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		idx := 0
		currentWord := ""
		for _, token := range startingWord {
			letter := string(token)
			currentWord += letter
			if idx < 2 {
				continue
			}
			if strings.HasPrefix(scanner.Text(), letter) {
				usernames += strings.Join([]string{currentWord, scanner.Text()}, "")
				usernames += "\n"
			}
			idx++
		}
	}
	return usernames, nil
}
