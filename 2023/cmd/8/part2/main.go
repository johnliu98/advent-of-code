package main

import (
	"fmt"
	"regexp"

	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

const (
	left  byte = byte('L')
	right byte = byte('R')
	start byte = byte('A')
	goal  byte = byte('Z')
)

func main() {
	lines := read.ReadLines("./cmd/8/input.txt")

	network := make(network)
	starts := make([]node, 0)
	re := regexp.MustCompile(`\w\w\w`)
	for _, line := range lines[2:] {
		matches := re.FindAll([]byte(line), 3)
		network[node(matches[0])] = [2]node{node(matches[1]), node(matches[2])}
		if n := node(matches[0]); n.start() {
			starts = append(starts, n)
		}
	}

	directions := lines[0]
	cycles := make([]int, len(starts))
	for i := range starts {
		var count int
		for n := starts[i]; ; {
			d := directions[count%len(directions)]
			count++

			n = network.step(n, d)

			if n.goal() {
				cycles[i] = count
				break
			}
		}
	}

	ans := leastCommonMultiple(cycles...)

	fmt.Println("Answer: ", ans)
}

type network map[node][2]node

func (n network) step(node node, d byte) node {
	if d == left {
		return n[node][0]
	}
	return n[node][1]
}

type node string

func (n node) start() bool {
	return n[2] == start
}

func (n node) goal() bool {
	return n[2] == goal
}

func leastCommonMultiple(ints ...int) int {
	if len(ints) == 0 {
		return 0
	}

	if len(ints) == 1 {
		return ints[0]
	}

	result := ints[0] * ints[1] / greatestCommonDivisor(ints[0], ints[1])

	for i := 2; i < len(ints); i++ {
		result = leastCommonMultiple(result, ints[i])
	}

	return result
}

func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
