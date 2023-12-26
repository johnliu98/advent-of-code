package main

import (
	"fmt"

	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

func main() {
	input := read.ReadLines("input.txt")

	tiles := make(tiles, len(input))
	var startIdx index
	for i, row := range input {
		tiles[i] = make([]tile, len(row))
		for j, r := range row {
			t := tile(r)
			tiles[i][j] = t
			if t == start {
				startIdx = index{i, j}
			}
		}
	}

	var search searchIndex
	for _, d := range []direction{left, right, up, down} {
		search = tiles.next(searchIndex{i: startIdx, dir: d})
		if search.dir != none {
			break
		}
	}

	visited := []index{startIdx}
	for search.i != startIdx {
		visited = append(visited, search.i)
		search = tiles.next(search)
	}
	visited = append(visited, search.i)

	area := shoelaceFormula(visited)

	ans := picksTheorem(area, len(visited)-1)

	fmt.Println("Answer: ", ans)
}

func picksTheorem(a, b int) int {
	fmt.Println("b:", b)
	return a - b/2 + 1
}

func shoelaceFormula(xy []index) int {
	area := 0
	for i := 0; i < len(xy)-1; i++ {
		area += (xy[i][0] * xy[i+1][1]) - (xy[i+1][0] * xy[i][1])
	}
	area /= 2
	if area < 0 {
		area = -area
	}
	return area
}

type searchIndex struct {
	i   index
	dir direction
}

type index [2]int

func (i index) move(d direction) index {
	s := d.step()
	return index{
		i[0] + s[0],
		i[1] + s[1],
	}
}

type tiles [][]tile

func (t tiles) get(i index) tile {
	if i[0] < 0 || i[0] >= len(t) {
		return 0
	}
	row := t[i[0]]
	if i[1] < 0 || i[1] >= len(row) {
		return 0
	}
	return tile(t[i[0]][i[1]])
}

func (t tiles) next(s searchIndex) searchIndex {
	var n searchIndex
	n.i = s.i.move(s.dir)
	dirs := t.get(n.i).directions()
	if dirs == nil {
		return n
	}
	if s.dir.opposite(dirs[0]) {
		n.dir = dirs[1]
	}
	if s.dir.opposite(dirs[1]) {
		n.dir = dirs[0]
	}
	return n
}

type tile rune

const (
	northSouth tile = '|'
	eastWest   tile = '-'
	northEast  tile = 'L'
	northWest  tile = 'J'
	southWest  tile = '7'
	southEast  tile = 'F'
	ground     tile = '.'
	start      tile = 'S'
)

func (t tile) directions() []direction {
	switch t {
	case northSouth:
		return []direction{up, down}
	case eastWest:
		return []direction{left, right}
	case northEast:
		return []direction{up, right}
	case northWest:
		return []direction{up, left}
	case southWest:
		return []direction{left, down}
	case southEast:
		return []direction{down, right}
	default:
		return nil
	}
}

type direction int

const (
	none direction = iota
	left
	right
	up
	down
)

func (d direction) step() index {
	switch d {
	case left:
		return index{0, -1}
	case right:
		return index{0, 1}
	case up:
		return index{-1, 0}
	case down:
		return index{1, 0}
	default:
		return index{}
	}
}

func (d direction) opposite(x direction) bool {
	return d.step().move(x) == index{}
}
