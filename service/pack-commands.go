package service

func PackCommands() string {
	commandsMessage := "Available commands:\n" +
		"/generategame startingWord(s) generates a game-oriented username\n" +
		"/generatedate startingWord(s) generates a date app-oriented username\n" +
		"/generateenterprise startingWord(s) generates an enterprise-oriented username\n" +
		"/generatemix generates a mix of all the above\n" +
		"/info shows the info for using this bot\n" +
		"/help shows some help\n" +
		"/commands shows available commands (this message)\n"
	return commandsMessage
}
