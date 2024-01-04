package main

import (
	"fmt"

	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

func main() {
	grid := newGrid(read.ReadLines("input.txt"))

	start := index{i: 0, j: 1}
	goal := index{i: len(grid) - 1, j: len(grid[0]) - 2}
	ans := grid.longestPath(start, goal)

	fmt.Println("Answer: ", ans)
}

type cell struct {
	index index
	tile  rune
	dist  int
}

type grid [][]*cell

func newGrid(cells []string) grid {
	grid := make(grid, len(cells))
	for i, row := range cells {
		grid[i] = make([]*cell, len(row))
		for j, c := range row {
			grid[i][j] = &cell{
				index: index{i: i, j: j},
				tile:  c,
			}
		}
	}
	return grid
}

func (g grid) longestPath(startIndex, goalIndex index) int {
	var longest int

	parent := make(map[*cell]*cell)

	start := g[startIndex.i][startIndex.j]

	queue := []*cell{start}
	inQueue := make(map[*cell]struct{})

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		delete(inQueue, current)

		if current.index == goalIndex {
			longest = max(longest, current.dist)
		}

		for _, d := range []direction{up, left, down, right} {
			if current.tile == '^' && d != up {
				continue
			}
			if current.tile == '<' && d != left {
				continue
			}
			if current.tile == 'v' && d != down {
				continue
			}
			if current.tile == '>' && d != right {
				continue
			}

			index := current.index.move(d)
			if g.outOfBounds(index) {
				continue
			}

			neighbor := g.get(index)

			if neighbor.tile == '#' {
				continue
			}

			if neighbor == parent[current] {
				continue
			}

			dist := current.dist + 1
			if dist < neighbor.dist {
				continue
			}

			neighbor.dist = dist
			parent[neighbor] = current

			if _, ok := inQueue[neighbor]; !ok {
				queue = append(queue, neighbor)
				inQueue[neighbor] = struct{}{}
			}
		}
	}

	return longest
}

func (g grid) get(i index) *cell {
	return g[i.i][i.j]
}

func (g grid) outOfBounds(i index) bool {
	if i.i < 0 {
		return true
	}
	if i.i >= len(g) {
		return true
	}
	if i.j < 0 {
		return true
	}
	if i.j >= len(g[0]) {
		return true
	}
	return false
}
