package lecture3

func sendToChannel(text string, messages chan string) {
	messages <- text
}
