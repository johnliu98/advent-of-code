package main

import (
	"container/heap"
	"fmt"

	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

// Input size is 141x141. Max heat loss per cell is 9.
// Use 200000 > 141*141*9 instead of math.MaxInt64 so avoid overlow.
const maxHeatLoss = 200000

func main() {
	grid := newGrid(read.ReadLines("input.txt"))

	start := index{i: 0, j: 0}
	goal := index{i: len(grid) - 1, j: len(grid[0]) - 1}
	ans := grid.minimumHeatLoss(start, goal)

	fmt.Println("Answer: ", ans)
}

type cell struct {
	index     index
	state     state
	loss      int
	gLoss     int
	fLoss     int
	heapIndex int
}

type grid [][][]*cell

func newGrid(cells []string) grid {
	grid := make(grid, len(cells))
	for i, row := range cells {
		grid[i] = make([][]*cell, len(row))
		for j, c := range row {
			grid[i][j] = make([]*cell, numStates)
			for k := 0; k < numStates; k++ {
				grid[i][j][k] = &cell{
					index: index{i: i, j: j},
					state: state(k),
					loss:  int(c - '0'),
					gLoss: maxHeatLoss,
					fLoss: maxHeatLoss,
				}
			}
		}
	}
	return grid
}

func (g grid) minimumHeatLoss(startIndex, goalIndex index) int {
	pq := make(priorityQueue, 0)
	heap.Init(&pq)

	open := make(map[*cell]struct{})
	closed := make(map[*cell]struct{})

	parent := make(map[*cell]*cell)

	start := g[startIndex.i][startIndex.j][none]
	start.gLoss = 0
	start.fLoss = goalIndex.dist(startIndex)

	heap.Push(&pq, start)
	open[start] = struct{}{}
	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*cell)
		closed[current] = struct{}{}
		if current.index == goalIndex {
			return current.gLoss
		}

		for _, d := range []direction{up, left, down, right} {
			if current.state.direction().opposite(d) {
				continue
			}

			index := current.index.move(d)
			if g.outOfBounds(index) {
				continue
			}

			state := current.state.move(d)
			if state.invalid() {
				continue
			}

			neighbor := g.get(index, state)

			if _, ok := closed[neighbor]; ok {
				continue
			}

			gLoss := current.gLoss + neighbor.loss
			if gLoss >= neighbor.gLoss {
				continue
			}

			neighbor.gLoss = gLoss
			neighbor.fLoss = gLoss + goalIndex.dist(neighbor.index)
			heap.Fix(&pq, neighbor.heapIndex)

			parent[neighbor] = current

			if _, ok := open[neighbor]; !ok {
				heap.Push(&pq, neighbor)
				open[neighbor] = struct{}{}
			}
		}
	}

	panic("cound not find goal")
}

func (g grid) get(i index, s state) *cell {
	return g[i.i][i.j][s]
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
