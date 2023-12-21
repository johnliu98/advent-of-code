package read

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

func GameID(s string) int {
	x := strings.Split(s, ":")
	x = strings.Split(x[0], " ")
	i, err := strconv.Atoi(x[1])
	if err != nil {
		panic(err)
	}
	return i
}

func Sets(s string) []string {
	x := strings.Split(s, ":")
	return strings.Split(x[1], ";")
}
