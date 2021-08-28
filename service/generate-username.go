package service

import (
	"strings"
)

func GenerateUsername(startingWord string) string {
	var usernames string
	newline := "\n"

	for _, char := range startingWord {
		charS := string(char)
		if charS == "a" || charS == "e" || charS == "i" || charS == "o" || charS == "u" {
			usernames += strings.Replace(startingWord, charS, "x", 1) + newline
		}
	}

	return usernames
}
