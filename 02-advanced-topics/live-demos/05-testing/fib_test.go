package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
	n := 10
	expected := 55
	actual := Fibonacci(n)

	if expected != actual {
		t.Fatalf("expected %d, got %d", expected, actual)
	}
}

func TestFibMultiple(t *testing.T) {
	tests := []struct{ seqNo, expected int }{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{5, 5},
		{10, 55},
		{20, 6765},
		{30, 832040},
		{40, 102334155},
		{45, 1134903170},
		{45, 1134903170},
	}

	for _, v := range tests {
		n := v.seqNo
		expected := v.expected

		t.Run(fmt.Sprintf("n = %d", n), func(t *testing.T) {
			t.Parallel()
			actual := Fibonacci(n)

			assert.Equal(t, actual, expected)
		})
	}
}
