package service

func PackCommands() string {
	commandsMessage := "Available commands:\n" +
		"/generategame `number` startingWord(s) generates `number` game-oriented username\n" +
		"/generatedate `number` startingWord(s) generates `number` date app-oriented username\n" +
		"/generateenterprise `number` startingWord(s) generates `number` enterprise-oriented username\n" +
		"/generatemix `number` generates `number` mix of all the above\n" +
		"/info shows the info for using this bot\n" +
		"/help shows some help\n" +
		"/commands shows available commands (this message)\n"
	return commandsMessage
}
