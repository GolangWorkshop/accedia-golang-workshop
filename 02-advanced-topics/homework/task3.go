package main

import "fmt"

func main() {

	a := []int{1, 2, 3, 4}

	fmt.Println(a)

	b := filter[int](a, func(x int) bool {
		return x%2 == 0
	})

	fmt.Println(b)

	c := mapper[int, int](b, func(x int) int {
		return x * x
	})

	fmt.Println(c)
}

func filter[T any](slice []T, predicate func(T) bool) []T {
	var n []T
	for _, v := range slice {
		if predicate(v) {
			n = append(n, v)
		}
	}

	return n
}

func mapper[T any, V any](slice []T, delegate func(T) V) []V {
	var n []V
	for _, v := range slice {
		n = append(n, delegate(v))
	}

	return n
}
