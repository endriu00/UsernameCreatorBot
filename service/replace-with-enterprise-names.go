package service

import (
	"bufio"
	"os"
	"strings"
)

func ReplaceWithEnterpriseNames(startingWord string, repNum int) (string, error) {
	var usernames string
	var hits int = 0
	file, err := os.Open("names/enterprise.txt")
	if err != nil {
		return usernames, err
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if hits == repNum {
			break
		}
		match := strings.Join([]string{startingWord, scanner.Text()}, "")
		usernames = strings.Join([]string{usernames, match}, "\n")
		hits++
	}
	return usernames, nil
}
