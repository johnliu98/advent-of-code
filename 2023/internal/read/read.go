package read

import (
	"bufio"
	"os"
)

func ReadLines(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)

	lines := make([]string, 0)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	return lines
}
