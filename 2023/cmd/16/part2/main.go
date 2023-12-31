package main

import (
	"fmt"

	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

func main() {
	grid := grid{
		tiles: read.ReadLines("input.txt"),
		next:  make(map[beam][]beam),
	}

	var ans int
	for i := range grid.tiles {
		ans = max(ans, grid.numEnergized(beam{
			index:     index{i: i, j: 0},
			direction: right,
		}))
		ans = max(ans, grid.numEnergized(beam{
			index:     index{i: i, j: len(grid.tiles[0]) - 1},
			direction: left,
		}))
	}
	for j := range grid.tiles[0] {
		ans = max(ans, grid.numEnergized(beam{
			index:     index{i: 0, j: j},
			direction: down,
		}))
		ans = max(ans, grid.numEnergized(beam{
			index:     index{i: len(grid.tiles) - 1, j: j},
			direction: up,
		}))
	}

	fmt.Println("Answer: ", ans)
}

type grid struct {
	tiles     []string
	energized map[index]struct{}
	visited   map[beam]struct{}
	next      map[beam][]beam
}

func (g grid) numEnergized(b beam) int {
	g.energized = make(map[index]struct{})
	g.visited = make(map[beam]struct{})
	g.sendBeam(b)
	return len(g.energized)
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

	if next, ok := g.next[b]; ok {
		for _, n := range next {
			g.sendBeam(n)
		}
	}

	var next []beam
	switch g.get(b.index) {
	case '|':
		if b.direction == left || b.direction == right {
			next = append(next, b.move(up), b.move(down))
		} else {
			next = append(next, b.move(b.direction))
		}
	case '-':
		if b.direction == up || b.direction == down {
			next = append(next, b.move(left), b.move(right))
		} else {
			next = append(next, b.move(b.direction))
		}
	case '/':
		if b.direction == up {
			next = append(next, b.move(right))
		} else if b.direction == left {
			next = append(next, b.move(down))
		} else if b.direction == down {
			next = append(next, b.move(left))
		} else if b.direction == right {
			next = append(next, b.move(up))
		}
	case '\\':
		if b.direction == up {
			next = append(next, b.move(left))
		} else if b.direction == left {
			next = append(next, b.move(up))
		} else if b.direction == down {
			next = append(next, b.move(right))
		} else if b.direction == right {
			next = append(next, b.move(down))
		}
	default:
		next = append(next, b.move(b.direction))
	}

	g.next[b] = next
	for _, n := range next {
		g.sendBeam(n)
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
