package main

import (
	"fmt"

	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

func main() {
	grid := grid{
		tiles:     read.ReadLines("input.txt"),
		energized: make(map[index]struct{}),
		visited:   make(map[beam]struct{}),
	}
	start := beam{
		index:     index{i: 0, j: 0},
		direction: right,
	}

	ans := numEnergized(grid, start)

	fmt.Println("Answer: ", ans)
}

func numEnergized(g grid, b beam) int {
	g.sendBeam(b)
	return len(g.energized)
}

type grid struct {
	tiles     []string
	energized map[index]struct{}
	visited   map[beam]struct{}
}

func (g grid) sendBeam(b beam) {
	if g.outOfBounds(b.index) {
		return
	}
	if _, ok := g.visited[b]; ok {
		return
	}
	g.visited[b] = struct{}{}

	g.energized[b.index] = struct{}{}

	switch g.get(b.index) {
	case '|':
		if b.direction == left || b.direction == right {
			g.sendBeam(b.move(up))
			g.sendBeam(b.move(down))
		} else {
			g.sendBeam(b.move(b.direction))
		}
	case '-':
		if b.direction == up || b.direction == down {
			g.sendBeam(b.move(left))
			g.sendBeam(b.move(right))
		} else {
			g.sendBeam(b.move(b.direction))
		}
	case '/':
		if b.direction == up {
			g.sendBeam(b.move(right))
		} else if b.direction == left {
			g.sendBeam(b.move(down))
		} else if b.direction == down {
			g.sendBeam(b.move(left))
		} else if b.direction == right {
			g.sendBeam(b.move(up))
		}
	case '\\':
		if b.direction == up {
			g.sendBeam(b.move(left))
		} else if b.direction == left {
			g.sendBeam(b.move(up))
		} else if b.direction == down {
			g.sendBeam(b.move(right))
		} else if b.direction == right {
			g.sendBeam(b.move(down))
		}
	default:
		g.sendBeam(b.move(b.direction))
	}
}

func (g grid) get(i index) byte {
	return g.tiles[i.i][i.j]
}

func (g grid) outOfBounds(i index) bool {
	if i.i < 0 {
		return true
	}
	if i.i >= len(g.tiles) {
		return true
	}
	if i.j < 0 {
		return true
	}
	if i.j >= len(g.tiles[0]) {
		return true
	}
	return false
}

type beam struct {
	index     index
	direction direction
}

func (b beam) move(d direction) beam {
	switch d {
	case up:
		return beam{
			index:     index{b.index.i - 1, b.index.j},
			direction: d,
		}
	case left:
		return beam{
			index:     index{b.index.i, b.index.j - 1},
			direction: d,
		}
	case down:
		return beam{
			index:     index{b.index.i + 1, b.index.j},
			direction: d,
		}
	case right:
		return beam{
			index:     index{b.index.i, b.index.j + 1},
			direction: d,
		}
	default:
		panic("invalid direction")
	}
}

type index struct {
	i, j int
}

type direction int

const (
	up direction = iota
	left
	down
	right
)
