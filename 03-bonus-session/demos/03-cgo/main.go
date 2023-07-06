// main.go
package main

/*
  #include "hello.c"
  #include "sum.c"
*/
import "C"
import (
	"errors"
	"fmt"
	"log"
)

func main() {
	//Call to void function without params
	err := Hello()
	if err != nil {
		log.Fatal(err)
	}

	//Call to int function with two params
	res, err := makeSum(5, 4)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sum of 5 + 4 is %d\n", res)
}

func Hello() error {
	_, err := C.Hello() //We ignore first result as it is a void function
	if err != nil {
		return errors.New("error calling Hello function: " + err.Error())
	}

	return nil
}

func makeSum(a, b int) (int, error) {
	//Convert Go ints to C ints
	aC := C.int(a)
	bC := C.int(b)

	sum, err := C.sum(aC, bC)
	if err != nil {
		return 0, errors.New("error calling Sum function: " + err.Error())
	}

	//Convert C.int result to Go int
	res := int(sum)

	return res, nil
}
