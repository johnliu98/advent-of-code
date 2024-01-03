package main

import (
	"fmt"

	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

const numSteps = 64

func main() {
	g := read.ReadLines("input.txt")

	var start index
	rocks := make(map[index]struct{})
	for i, row := range g {
		for j, c := range row {
			if c == 'S' {
				start = index{i: i, j: j}
			}
			if c == '#' {
				rocks[index{i: i, j: j}] = struct{}{}
			}
		}
	}

	tiles := make(map[index]struct{})
	tiles[start] = struct{}{}
	for s := 0; s < numSteps; s++ {
		nextTiles := make(map[index]struct{})
		for i := range tiles {
			for _, d := range []direction{up, left, down, right} {
				next := i.move(d)
				if _, ok := rocks[next]; ok {
					continue
				}

				nextTiles[next] = struct{}{}
			}
		}
		tiles = nextTiles
	}

	ans := len(tiles)

	fmt.Println("Answer: ", ans)
}
