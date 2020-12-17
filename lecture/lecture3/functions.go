package lecture3

func SendToChannel(text string, messages chan string) {
	messages <- text
}
