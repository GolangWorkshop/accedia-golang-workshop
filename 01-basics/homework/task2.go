package main

import (
	"fmt"
	"strings"
)

func isLetter(c byte) bool {
	return ('a' <= c && c <= 'z')
}

func Task_2_solution() {
	input := "Z!y X"
	input = strings.ToLower(input)
	var lettersFromInput []byte
	for i := len(input) - 1; i >= 0; i-- {
		if isLetter(input[i]) {
			lettersFromInput = append(lettersFromInput, input[i]-'a'+1)
		}
	}
	for _, value := range lettersFromInput {
		fmt.Print(value)
	}
}
