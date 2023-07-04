package main

import (
	"fmt"
	"strings"
)

func uncensor(censored string, uncensored string) string {
	censoredParts := strings.Split(censored, "*")
	uncensoredRunes := []rune(uncensored)

	var result strings.Builder

	for i := 0; i < len(censoredParts)-1; i++ {
		result.WriteString(censoredParts[i])

		if len(uncensoredRunes) > i {
			result.WriteRune(uncensoredRunes[i])
		}
	}

	if len(censoredParts) > 0 {
		result.WriteString(censoredParts[len(censoredParts)-1])
	}

	return result.String()
}

func Task_5_solution() {

	input := "*h*s *s v*ry *tr*ng*"
	missingLetters := "Tiiesae"
	fmt.Println(uncensor(input, missingLetters))

}
