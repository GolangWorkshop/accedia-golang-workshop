package main

import (
	"fmt"
	"time"
)

func display(message string) {

	fmt.Println(message)
}

func displayTimeOut(message string) {

	// infinite for loop
	for {
		fmt.Println(message)

		// to sleep the process for 1 second
		time.Sleep(time.Second * 1)
	}
}

func main() {

	// call goroutine
	go displayTimeOut("Additional goroutine")
	go displayTimeOut("Process 3")
	go displayTimeOut("Process 4")
	//go display("Additional goroutine 1")

	//function call
	//display("Text from main goroutine")
	displayTimeOut("Text from main goroutine")
}
