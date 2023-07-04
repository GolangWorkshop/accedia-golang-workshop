package main

import (
	"fmt"
	"strconv"
	"strings"
)

func splitNumbers(s []int) (int, int) {
	sumOfNeg, sumOfPos := 0, 0
	for _, n := range s {
		if n < 0 {
			sumOfNeg += n
		} else {
			sumOfPos += n
		}
	}
	return sumOfNeg, sumOfPos
}

func Task_3_solution() {

	input := "10 18 30 0"
	numbers := strings.Fields(input)
	var allNumbers []int
	for _, value := range numbers {
		if num, err := strconv.Atoi(value); err == nil {
			if num != 0 {
				allNumbers = append(allNumbers, num)
			} else {
				break
			}
		} else {
			fmt.Println(err)
		}
	}

	sumNeg, sumPos := splitNumbers(allNumbers)
	if -sumNeg == sumPos {
		fmt.Println("Yes")
	} else if diff := sumPos - (-sumNeg); diff < 0 {
		fmt.Println(-diff)
	} else {
		fmt.Println(diff)
	}
}
