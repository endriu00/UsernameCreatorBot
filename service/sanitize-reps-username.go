package service

import (
	"strconv"
	"strings"
	"text/scanner"
)

/*
SanitizeRepsUsername is a simple function that returns:
- the starting word for the username generation
- the number of usernames to generate
- an error
*/
func SanitizeRepsUsername(rawCommand string) (string, int, error) {
	var stringScan scanner.Scanner
	//Initialize scanner for the string rawCommand, which has the form int, string,
	//where int is the number of reps representing the iterations of the generate commands
	//and string is the string from which the algorithm starts for generating new usernames.
	stringScan.Init(strings.NewReader(rawCommand))

	//Scan next token. Expected token: a number
	_ = stringScan.Scan()

	//Convert the processed token into an int. If there is an error, return it.
	reps, err := strconv.Atoi(stringScan.TokenText())
	if err != nil {
		return "", 0, err
	}

	//Scan next token. Expected token: a string
	_ = stringScan.Scan()

	//Get the last token, the username
	startingWord := stringScan.TokenText()

	return startingWord, reps, nil
}
