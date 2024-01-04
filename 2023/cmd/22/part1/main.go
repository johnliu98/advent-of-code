package main

import (
	"fmt"
	"sort"

	"github.com/johnliu98/advent-of-code/2023/internal/parse"
	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

func main() {
	lines := read.ReadLines("input.txt")
	bricks := make(bricks, len(lines))
	for i, line := range lines {
		ints := parse.Ints(line)
		bricks[i] = brick{
			a: point{x: ints[0], y: ints[1], z: ints[2]},
			b: point{x: ints[3], y: ints[4], z: ints[5]},
		}
	}

	sort.Sort(bricks)

	var height grid
	var piece grid
	crucial := make(map[int]struct{})
	for i, b := range bricks {
		var h int
		for x := b.a.x; x <= b.b.x; x++ {
			for y := b.a.y; y <= b.b.y; y++ {
				h = max(h, height[x][y])
			}
		}

		support := make(map[int]struct{})
		for x := b.a.x; x <= b.b.x; x++ {
			for y := b.a.y; y <= b.b.y; y++ {
				if height[x][y] == h && piece[x][y] != 0 {
					support[piece[x][y]] = struct{}{}
				}
			}
		}

		if len(support) == 1 {
			for p := range support {
				crucial[p] = struct{}{}
			}
		}

		l := (b.b.z - b.a.z + 1)
		for x := b.a.x; x <= b.b.x; x++ {
			for y := b.a.y; y <= b.b.y; y++ {
				height[x][y] = h + l
				piece[x][y] = i + 1
			}
		}
	}

	ans := len(bricks) - len(crucial)

	fmt.Println("Answer: ", ans)
}

const gridSize = 10

type grid [gridSize][gridSize]int

type point struct {
	x, y, z int
}

type brick struct {
	a, b point
}

type bricks []brick

func (b bricks) Len() int {
	return len(b)
}

func (b bricks) Less(i, j int) bool {
	return b[i].b.z < b[j].b.z
}

func (b bricks) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
