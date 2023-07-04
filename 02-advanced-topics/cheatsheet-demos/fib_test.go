package main

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	seqNo := 1
	expected := 1
	actual := fibonacci(seqNo)

	if expected != actual {
		t.Errorf("fibonacci(%d) = %d, got %d", seqNo, expected, actual)
	}
}

func TestFibMultiple(t *testing.T) {
	cases := []struct{ seqNo, expected int }{
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
		// {45, 1134903170}, // will take about double the time
	}
	for _, tc := range cases {
		actual := fibonacci(tc.seqNo)

		if tc.expected != actual {
			t.Errorf("fibonacci(%d) = %d, got %d", tc.seqNo, tc.expected, actual)
		}
	}
}

func TestFibMultipleParallel(t *testing.T) {
	cases := []struct{ seqNo, expected int }{
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
		{45, 1134903170},
		{45, 1134903170},
	}
	for n, tc := range cases {
		testCase := tc

		t.Run(fmt.Sprintf("test case %d", n), func(t *testing.T) {
			t.Parallel()
			actual := fibonacci(testCase.seqNo)

			if testCase.expected != actual {
				t.Errorf("fibonacci(%d) = %d, got %d", testCase.seqNo, testCase.expected, actual)
			}
		})
	}
}
