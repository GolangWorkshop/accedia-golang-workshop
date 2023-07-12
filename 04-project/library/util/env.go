package util

import (
	"bufio"
	"os"
	"strings"
)

func LoadEnvFromFile(file string) (bool, error) {
	f, err := os.Open(file)
	if err != nil {
		return false, err
	}

	s := bufio.NewScanner(f)

	s.Split(bufio.ScanLines)

	for s.Scan() {
		line := s.Text()

		i := strings.Index(line, "=")

		if i < 0 {
			continue
		}

		name := line[0:i]
		val := line[i+1:]

		os.Setenv(name, val)
	}

	return true, nil
}
