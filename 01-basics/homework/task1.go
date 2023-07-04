package main

import (
	"fmt"
	"strconv"
	"strings"
)

func calculateFactorial(num int) int {
	if num <= 1 {
		return 1
	} else {
		return num * calculateFactorial(num-1)
	}
}

func calculateDigits(num int) int {
	if num == 0 {
		return 1
	}
	count := 0
	for num != 0 {
		num /= 10
		count++
	}
	return count
}

func Task_1_solution() {
	input := "1 2 3 4 5"
	numbers := strings.Fields(input)
	for _, value := range numbers {
		if num, err := strconv.Atoi(value); err == nil {
			fmt.Print(calculateDigits(calculateFactorial(num)), " ")
		} else {
			fmt.Println(err)
		}
	}

}
