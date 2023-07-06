package main

import (
	"fmt"
	"unsafe"
)

type Employee struct {
	ID     int
	Name   string
	Age    int16
	Gender string
	Active bool
}

//Employee struct size = 8(ID) + 16(Name) + 2(Age) + 16(Gender) + 1(Active) = 43 bytes

func main() {
	var e Employee
	fmt.Printf("Size of %T struct: %d bytes", e, unsafe.Sizeof(e))
}

//optimized Employee
//type Employee struct {
//	Name string
//	Gender string
//	ID int
//	Age int16
//	Active bool
//}
