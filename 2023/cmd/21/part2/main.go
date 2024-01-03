package main

import (
	"fmt"

	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

// 26501365 = 202300 * 131 + 65

func main() {
	var numPlots []int
	const numSteps = 3
	for steps := 0; steps < numSteps; steps++ {
		numSteps := steps*131 + 65

		g := newGarden("input.txt")

		if numSteps%2 == 0 {
			g.frontier[g.start] = struct{}{}
			g.plots[g.start] = struct{}{}
		} else {
			for _, d := range []direction{up, left, down, right} {
				ii := g.start.move(d)
				if g.rock(ii) || g.visited(ii) {
					continue
				}
				g.frontier[ii] = struct{}{}
				g.plots[ii] = struct{}{}
			}
		}

		for s := 0; s < numSteps-1; s += 2 {
			nextFrontier := make(map[index]struct{})
			for i := range g.frontier {
				for _, d := range []direction{up, left, down, right} {
					ii := i.move(d)
					if g.rock(ii) || g.front(ii) {
						continue
					}
					for _, dd := range []direction{up, left, down, right} {
						iii := ii.move(dd)
						if g.rock(iii) || g.visited(iii) {
							continue
						}
						nextFrontier[iii] = struct{}{}
						g.plots[iii] = struct{}{}
					}
				}
			}
			g.frontier = nextFrontier
		}

		numPlots = append(numPlots, len(g.plots))
	}

	extrapolation := extrapolate(numPlots, 202301-numSteps)

	ans := extrapolation[len(extrapolation)-1]

	fmt.Println("Answer: ", ans)
}

func extrapolate(values []int, count int) []int {
	if allZeros(values) {
		return make([]int, count)
	}

	diff := make([]int, len(values)-1)
	for i := 0; i < len(values)-1; i++ {
		diff[i] = values[i+1] - values[i]
	}

	diffExt := extrapolate(diff, count)

	for i := len(diffExt) - count; i < len(diffExt); i++ {
		v := values[len(values)-1] + diffExt[i]
		values = append(values, v)
	}

	return values
}

func allZeros(values []int) bool {
	for _, v := range values {
		if v != 0 {
			return false
		}
	}
	return true
}

type garden struct {
	tiles    []string
	start    index
	rocks    map[index]struct{}
	plots    map[index]struct{}
	frontier map[index]struct{}
	max      index
}

func newGarden(input string) garden {
	var g garden

	g.tiles = read.ReadLines(input)
	g.plots = make(map[index]struct{})
	g.rocks = make(map[index]struct{})
	g.frontier = make(map[index]struct{})

	for i, row := range g.tiles {
		for j, c := range row {
			if c == 'S' {
				g.start = index{i: i, j: j}
			}
			if c == '#' {
				g.rocks[index{i: i, j: j}] = struct{}{}
			}
			if i > g.max.i {
				g.max.i = i
			}
			if j > g.max.j {
				g.max.j = j
			}
		}
	}

	return g
}

func (g garden) rock(i index) bool {
	_, isRock := g.rocks[g.wrap(i)]
	return isRock
}

func (g garden) visited(i index) bool {
	_, isVisited := g.plots[i]
	return isVisited
}

func (g garden) front(i index) bool {
	_, inFront := g.frontier[i]
	return inFront
}

func (g garden) wrap(i index) index {
	for i.i < 0 {
		i.i += g.max.i + 1
	}
	for i.i > g.max.i {
		i.i -= g.max.i + 1
	}
	for i.j < 0 {
		i.j += g.max.j + 1
	}
	for i.j > g.max.j {
		i.j -= g.max.j + 1
	}
	return i
}
