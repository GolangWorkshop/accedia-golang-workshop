package main

import (
	"bufio"
	"compress/zlib"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
)

func readLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	line, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.New("Could not read line.")
	}

	// remove \n from input string
	line = strings.Replace(line, "\n", "", -1)

	if len(line) == 0 {
		return "", errors.New("No input provided.")
	}

	return line, nil
}

func zlib_compress(input io.ReadCloser, outputFile string) {
	// get a handle for compressed file
	compressedFileWriter, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer compressedFileWriter.Close()

	// get a zlib compression writer that will write to output
	zlibCompressionWriter, err := zlib.NewWriterLevel(compressedFileWriter, zlib.BestCompression)
	if err != nil {
		panic(err)
	}
	defer zlibCompressionWriter.Close()

	_, err = io.Copy(zlibCompressionWriter, input)
	if err != nil {
		panic(err)
	}
}

func zlib_decompress(inputFile string, outputFile string) {

	// get a handle for compressed file
	compressedFileReader, err := os.OpenFile(inputFile, os.O_RDONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer compressedFileReader.Close()

	// get a handle for decompressed file
	decompressedFileWriter, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer decompressedFileWriter.Close()

	r, err := zlib.NewReader(compressedFileReader)
	if err != nil {
		panic(err)
	}

	io.Copy(decompressedFileWriter, r)
	r.Close()
}

func main() {
	url, err := readLine()
	if err != nil {
		panic(err)
	}

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	zlib_compress(res.Body, "compressed.zlib")

	zlib_decompress("compressed.zlib", "decompressed.pdf")
}
