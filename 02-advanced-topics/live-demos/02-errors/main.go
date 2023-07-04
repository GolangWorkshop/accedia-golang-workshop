package main

import (
	"errors"
	"fmt"
)

type DivideByZeroError struct{}

func (e *DivideByZeroError) Error() string {
	return "Cannot divide by zero"
}

func main() {
	numerator := 10
	denominator := 0

	//Panic
	//divideByZero(numerator, denominator)

	//Panic and recover
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("Recovered from panic:", err)
	//	}
	//}()

	// Example scenario: dividing by zero
	//divideByZero(numerator, denominator)
	//fmt.Println("After divideByZero") // This line won't be reached because of the panic

	//Custom Error
	//result, err := divideWithError(numerator, denominator)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//} else {
	//	fmt.Println("Result:", result)
	//}

	//Wrap/Unwrap
	//result, err := divideWithError(numerator, denominator)
	//if err != nil {
	//	wrappedErr := fmt.Errorf("error occurred during division: %w", err)
	//	fmt.Println("Wrapped error:", wrappedErr)
	//
	//	// Unwrap the error to get the original error value
	//	originalErr := errors.Unwrap(wrappedErr)
	//	fmt.Println("Original error:", originalErr)
	//} else {
	//	fmt.Println("Result:", result)
	//}

	//Error Is / As
	result, err := divideWithError(numerator, denominator)
	if err != nil {
		// Check if the error is of type DivideByZeroError
		if errors.Is(err, &DivideByZeroError{}) {
			fmt.Println("Error: Cannot divide by zero")
		} else {
			fmt.Println("Error:", err)
		}

		// Try to extract the DivideByZeroError from the error chain
		var divideErr *DivideByZeroError
		if errors.As(err, &divideErr) {
			fmt.Println("DivideByZeroError found:", divideErr)
		}
	} else {
		fmt.Println("Result:", result)
	}
}

// Panic and recover
func divideByZero(numerator, denominator int) {

	result := numerator / denominator // This will cause a runtime panic
	fmt.Println("Result:", result)    // This line won't be executed
}

// Defined new error type
func divideWithError(numerator, denominator int) (int, error) {
	if denominator == 0 {
		//return 0, errors.New("cannot divide by zero")
		return 0, &DivideByZeroError{}
	}
	return numerator / denominator, nil
}
