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
	supportedBy := make(map[int][]int)
	supports := make(map[int][]int)
	for i, b := range bricks {
		var h int
		for x := b.a.x; x <= b.b.x; x++ {
			for y := b.a.y; y <= b.b.y; y++ {
				h = max(h, height[x][y])
			}
		}

		sb := make(map[int]struct{})
		for x := b.a.x; x <= b.b.x; x++ {
			for y := b.a.y; y <= b.b.y; y++ {
				if height[x][y] == h && piece[x][y] != 0 {
					sb[piece[x][y]] = struct{}{}
				}
			}
		}

		for p := range sb {
			supportedBy[i+1] = append(supportedBy[i+1], p)
			supports[p] = append(supports[p], i+1)
		}

		l := (b.b.z - b.a.z + 1)
		for x := b.a.x; x <= b.b.x; x++ {
			for y := b.a.y; y <= b.b.y; y++ {
				height[x][y] = h + l
				piece[x][y] = i + 1
			}
		}
	}

	var ans int
	visited := make(map[int]struct{})
	for _, ps := range supportedBy {
		if len(ps) > 1 {
			continue
		}

		if _, ok := visited[ps[0]]; ok {
			continue
		}
		visited[ps[0]] = struct{}{}

		queue := []int{ps[0]}
		fallen := make(map[int]struct{})
		for len(queue) > 0 {
			p := queue[0]

			queue = queue[1:]
			fallen[p] = struct{}{}

			for _, n := range supports[p] {
				var supported bool
				for _, s := range supportedBy[n] {
					if _, ok := fallen[s]; !ok {
						supported = true
						break
					}
				}

				if !supported {
					queue = append(queue, n)
					fallen[n] = struct{}{}
				}
			}
		}

		ans += len(fallen) - 1
	}

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
