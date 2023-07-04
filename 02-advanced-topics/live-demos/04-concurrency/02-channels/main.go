package main

import "fmt"

func sendMessage(message string, channel chan<- string) {
	// Send the message to the channel
	channel <- message
}

func main() {
	// Create a channel of type string
	messageChannel := make(chan string)

	// Start a goroutine to send a message
	go sendMessage("Win gopher!", messageChannel)

	// Receive the message from the channel
	receivedMessage := <-messageChannel

	// Print the received message
	fmt.Println(receivedMessage)
}
