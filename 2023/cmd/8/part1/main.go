package main

import (
	"fmt"
	"regexp"

	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

const (
	left  byte = byte('L')
	right byte = byte('R')
)

func main() {
	lines := read.ReadLines("./cmd/8/input.txt")

	network := make(map[string][2]string)
	re := regexp.MustCompile(`\w\w\w`)
	for _, line := range lines[2:] {
		matches := re.FindAll([]byte(line), 3)
		network[string(matches[0])] = [2]string{
			string(matches[1]),
			string(matches[2]),
		}
	}

	directions := lines[0]
	var count int
	const goal = "ZZZ"
	for curr := "AAA"; curr != goal; {
		i := count % len(directions)
		d := directions[i]
		count++

		switch d {
		case left:
			curr = network[curr][0]
		case right:
			curr = network[curr][1]
		default:
			panic(fmt.Sprintf("invalid direction: %s", string(d)))
		}
	}

	fmt.Println("Answer: ", count)
}
