package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	// localFileToStdout()
	// networkResponseToStdout()
	// writeToFile()
	completeExample()
}

func localFileToStdout() {
	f, err := os.Open("README.md") // opens a file for read only
	if err != nil {
		panic(err) // handle better in prod
	}
	defer f.Close() // do you remember what defer does?

	buf := make([]byte, 32) // 32b buffer, will store the data we read

	for {
		n, err := f.Read(buf) // fills the buffer with as much data as can be read
		if err == io.EOF {
			fmt.Println("EOF reached")
			break
		}
		// handle any other error we might get
		if err != nil {
			fmt.Println(err)
			break
		}
		if n > 0 {
			fmt.Println(string(buf[:n])) // print the bytes from 0 to n casted as string
		}
	}
}

func networkResponseToStdout() {
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		panic(err) // handle better in prod
	}
	defer conn.Close()

	conn.Write([]byte("GET / HTTP/1.1\r\n\r\n"))

	buf := make([]byte, 100)

	for {
		n, err := conn.Read(buf) // fills the buffer with as much data as can be read
		if err == io.EOF {
			fmt.Println("EOF reached")
			break
		}
		// handle any other error we might get
		if err != nil {
			fmt.Println(err)
			break
		}
		if n > 0 {
			fmt.Println(string(buf[:n])) // print the bytes from 0 to n casted as string
		}
	}
}

func writeToFile() {
	f, err := os.OpenFile("file.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600) // opens a file for read only
	if err != nil {
		panic(err)
	}

	defer f.Close() // do you remember what defer does?

	f.Write([]byte("write this data into the file\n"))
}

type counter struct {
	total uint64
}

// Write
func (c *counter) Write(b []byte) (int, error) {
	c.total += uint64(len(b)) // 32kb at a time
	progress := float64(c.total) / (1024 * 1024)
	fmt.Printf("\rDownloading %f MB...", progress)
	return len(b), nil
}

// big gzipped file
// http://storage.googleapis.com/books/ngrams/books/googlebooks-eng-all-5gram-20120701-0.gz

func completeExample() {
	res, err := http.Get("http://storage.googleapis.com/books/ngrams/books/googlebooks-eng-all-5gram-20120701-0.gz")
	if err != nil {
		panic(err)
	}
	// download the file into our local fs
	local, err := os.OpenFile("download-5gram.txt", os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer local.Close()

	dec, err := gzip.NewReader(res.Body)
	if err != nil {
		panic(err)
	}
	// copy res.Body into local file
	if _, err := io.Copy(local, io.TeeReader(dec, &counter{})); err != nil {
		panic(err)
	}
}
