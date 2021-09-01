package service

import (
	"bufio"
	"os"
	"strings"
)

func ReplaceWithDateNames(startingWord string, repNum int) (string, error) {
	var usernames string
	var hits int = 0
	file, err := os.Open("names/date.txt")
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if hits == repNum {
			break
		}
		if startingWord[0] == scanner.Text()[0] {
			match := strings.Join([]string{scanner.Text(), startingWord}, "")
			usernames = strings.Join([]string{usernames, match}, "\n")
			hits++
		}
	}
	return usernames, nil
}
