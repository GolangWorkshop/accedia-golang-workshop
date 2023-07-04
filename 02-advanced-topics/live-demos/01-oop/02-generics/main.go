package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type homeworkSolution interface {
	constraints.Integer | constraints.Float | string | []string
}

type homework[T homeworkSolution] struct {
	description string
	solution    T
}

func findMax[T constraints.Ordered](elements []T) (T, bool) {
	var max T

	if len(elements) < 1 {
		return max, false
	}

	if len(elements) == 1 {
		return elements[0], true
	}

	max = elements[0]

	for _, v := range elements[1:] {
		if max < v {
			max = v
		}
	}

	return max, true
}

func main() {
	h := homework[int]{
		description: "Whats the 10 fibbonacci number",
		solution:    55,
	}

	h2 := homework[string]{
		description: "Whats the 3rd planet from the Sun",
		solution:    "Earth",
	}

	arr := []int{1, 2, 3}
	arr2 := []float64{1.2, 2.3, 3.4}

	fmt.Println(findMax[int](arr))
	fmt.Println(findMax[float64](arr2))

	fmt.Println(h, h2)
}
