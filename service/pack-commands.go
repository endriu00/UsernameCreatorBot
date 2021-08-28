package service

func PackCommands() string {
	commandsMessage := "Available commands:\n" +
		"/generate\n" +
		"/info\n" +
		"/help\n" +
		"/commands\n"
	return commandsMessage
}
