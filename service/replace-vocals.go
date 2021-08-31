package service

import (
	"strings"
)

/*
ReplaceVocals is a function that replace the starting word, startingWord, with the string toReplace repNum times.
It is designed to replace vocals only.
*/
func ReplaceVocals(startingWord, toReplace string, repNum int) string {
	var usernames string
	var idx int = 0
	for _, char := range startingWord {
		if idx == repNum {
			break
		}
		charS := string(char)
		if charS == "a" || charS == "e" || charS == "i" || charS == "o" || charS == "u" {
			usernames += strings.Replace(startingWord, charS, toReplace, 1)
			usernames += "\n"
		}
		idx++
	}
	return usernames
}
