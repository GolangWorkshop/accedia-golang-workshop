package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const file = "test.txt"
const resource = "https://google.com/"

func main() {
	fromString("Super interesting string")
	// fromFile(file)
	// fromFileLineByLine(file)
	// toFile(file, "hello world")
	// fromNetworkToFile(file, resource)
}

func fromNetworkToFile(file string, resource string) {
	r, err := http.Get(resource)
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	written, err := io.Copy(f, r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("wrote ", written, " to file ", file)

}

func toFile(file string, data string) {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Write([]byte(data))
}

func fromFileLineByLine(file string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func fromFile(file string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b := make([]byte, 6)

	for {
		readBytes, err := f.Read(b)
		if err == io.EOF {
			fmt.Println("Whole file read")
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println(string(b[0:readBytes]))
	}
}

func fromString(str string) {
	r := strings.NewReader(str)

	b := make([]byte, 10)

	for {
		readBytes, err := r.Read(b)
		if err == io.EOF {
			fmt.Println("Whole file read")
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println(string(b[0:readBytes]))
	}
}

// check the readme for more examples
