package main

import "fmt"

func Task_4_solution() {

	input := "234256"

	// count occurrences of different digits
	var occurrences [10]int

	// remember last repeated digit
	lastRepeatedDigit := -1

	for _, unicodeValue := range input {
		// increment the counter of occurrences of the selected digit
		// removing the code of 0 returns the value of the rune of the number as int
		occurrences[int(unicodeValue-'0')]++
		// if the value exists: mark as last repeated digit
		if occurrences[int(unicodeValue-'0')] > 1 {
			lastRepeatedDigit = int(unicodeValue - '0')
		}
	}

	if lastRepeatedDigit != -1 {
		fmt.Println(lastRepeatedDigit)
	} else {
		fmt.Println("No")
	}
}
