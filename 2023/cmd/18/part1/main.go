package main

import (
	"fmt"
	"strings"

	"github.com/johnliu98/advent-of-code/2023/internal/conv"
	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

func main() {
	digPlan := read.ReadLines("input.txt")

	var area int
	var loop int
	current := trench{}
	for _, instruction := range digPlan {
		split := strings.Split(instruction, " ")
		d := direction(split[0][0])
		count := conv.IntFromString(split[1])

		next := current.move(d, count)

		area += current.i*next.j - next.i*current.j
		loop += count

		current = next
	}
	area /= 2
	if area < 0 {
		area = -area
	}

	interior := area - loop/2 + 1

	ans := loop + interior

	fmt.Println("Answer: ", ans)
}

type trench struct {
	i, j int
}

func (i trench) move(d direction, count int) trench {
	switch d {
	case up:
		return trench{i: i.i - count, j: i.j}
	case left:
		return trench{i: i.i, j: i.j - count}
	case down:
		return trench{i: i.i + count, j: i.j}
	case right:
		return trench{i: i.i, j: i.j + count}
	default:
		panic("invalid direction")
	}
}

type direction byte

const (
	up    direction = 'U'
	down  direction = 'D'
	left  direction = 'L'
	right direction = 'R'
)
