package service

func PackInfo(message Message) string {
	infoMessage := "Hey @" + message.From.Username + "\n" +
		"I am nomebot, your username creator bot!" +
		"To start with your username creation, just type\n" +
		"/generate `yourname`\n" +
		"where `yourname` is your name, your username starting point" +
		"or anything else.\n" +
		"Feel free to ask me for info again whenever you want with `/info` !"
	return infoMessage
}
