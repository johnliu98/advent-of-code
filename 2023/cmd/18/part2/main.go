package main

import (
	"fmt"

	"github.com/johnliu98/advent-of-code/2023/internal/parse"
	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

func main() {
	digPlan := read.ReadLines("input.txt")

	var area int
	var loop int
	var current trench
	for _, instruction := range digPlan {
		hexString, directionByte := parseInstruction(instruction)
		d := direction(directionByte - '0')
		steps := parse.Hex(hexString)

		next := current.move(d, steps)

		// Shoelace formula
		area += current.i*next.j - next.i*current.j
		loop += steps

		current = next
	}
	area /= 2
	if area < 0 {
		area = -area
	}

	// Pick's theorem
	interior := area - loop/2 + 1

	ans := loop + interior

	fmt.Println("Answer: ", ans)
}

func parseInstruction(s string) (string, byte) {
	var hash int
	for i, r := range s {
		if r == '#' {
			hash = i
			break
		}
	}
	return s[hash+1 : hash+6], s[hash+6]
}

type trench struct {
	i, j int
}

func (i trench) move(d direction, steps int) trench {
	switch d {
	case up:
		return trench{i: i.i - steps, j: i.j}
	case left:
		return trench{i: i.i, j: i.j - steps}
	case down:
		return trench{i: i.i + steps, j: i.j}
	case right:
		return trench{i: i.i, j: i.j + steps}
	default:
		panic("invalid direction")
	}
}

type direction byte

const (
	right direction = iota
	down
	left
	up
)
