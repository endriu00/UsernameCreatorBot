package service

func PackInfo(message Message) string {
	return "Hey @" + message.From.Username + "\n" +
		"I am UsernameGeneratorBot, your username creator bot!\n\n" +
		"To start with your username creation, search the menu for the available commands or just type\n\n`/commands`\n\n" +
		"This will show you a list of the available commands with a small description about what they do.\n" +
		"Here it is an example for a gaming username generation, just for you to familiarize with the commands:\n\n" +
		"/generategame `number` `yourname`\n" +
		"where:\n" +
		"`number` is how many username you want me to create\n" +
		"`yourname` is your name, your username starting point or anything else.\n" +
		"Note that you cannot ask for more than 15 usernames at once! You won't need them, don't worry.\n" +
		"Feel free to ask me for info again whenever you want with `/info` !"
}
