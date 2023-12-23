package main

import (
	"fmt"

	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

func main() {
	tiles := tiles(read.ReadLines("./cmd/10/input.txt"))

	var startIdx index
	for i, row := range tiles {
		for j, t := range row {
			if tile(t) == start {
				startIdx = index{i, j}
				break
			}
		}
	}

	var searchIndices [2]searchIndex
	var numFound int
	for _, d := range []direction{left, right, up, down} {
		i := startIdx.move(d)
		dirs := tiles.get(i).directions()
		if dirs == nil {
			continue
		}
		if d.opposite(dirs[0]) {
			searchIndices[numFound] = searchIndex{i: i, dir: dirs[1]}
			numFound++
		}
		if d.opposite(dirs[1]) {
			searchIndices[numFound] = searchIndex{i: i, dir: dirs[0]}
			numFound++
		}
		if numFound >= 2 {
			break
		}
	}

	count := 1
	for searchIndices[0].i != searchIndices[1].i {
		searchIndices[0] = tiles.next(searchIndices[0])
		searchIndices[1] = tiles.next(searchIndices[1])
		count++
	}

	fmt.Println("Answer: ", count)
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

type tiles []string

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
	left direction = iota
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
