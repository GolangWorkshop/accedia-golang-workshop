package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func number(number int) bool {
	if number < 2 {
		return false
	}
	sq_root := int(math.Sqrt(float64(number)))
	for i := 2; i <= sq_root; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func generateRandomNumbers(num int, randChannel chan<- int, errChannel chan<- error) {
	for i := 0; i < num; i++ {
		randomNumber := rand.Intn(100)
		if !number(randomNumber) {
			errChannel <- fmt.Errorf("random wrench number %d is not prime", randomNumber)
		} else {
			randChannel <- randomNumber
		}
	}
}

func Task_4_solution() {

	rand.Seed(time.Now().UnixNano())

	numbers := 20
	randChannel := make(chan int)
	errChannel := make(chan error)

	go generateRandomNumbers(numbers, randChannel, errChannel)

	sum := 0
	for i := 0; i < numbers; i++ {
		select {
		case randomNumber := <-randChannel:
			sum += randomNumber
		case err := <-errChannel:
			fmt.Printf("Error: %v\n", err)
		}
	}

	fmt.Printf("Sum of wrench numbers: %d\n", sum)
}
