package main

import (
	"fmt"

	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

func main() {
	image := image(read.ReadLines("./cmd/11/input.txt"))

	galaxies := image.findGalaxies()
	occupiedRows, occupiedCols := image.findOccupiedRowsAndCols(galaxies)

	var ans int
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			ans += galaxies[i].dist(galaxies[j], occupiedRows, occupiedCols)
		}
	}

	fmt.Println("Answer: ", ans)
}

type index [2]int

func (i index) dist(j index, occupiedRows, occupiedCols map[int]struct{}) int {
	const expansionDist = 1000000

	var rowDist int
	if j[0] < i[0] {
		i[0], j[0] = j[0], i[0]
	}
	for k := i[0]; k < j[0]; k++ {
		if _, ok := occupiedRows[k]; !ok {
			rowDist += expansionDist
		} else {
			rowDist++
		}
	}

	var colDist int
	if j[1] < i[1] {
		i[1], j[1] = j[1], i[1]
	}
	for k := i[1]; k < j[1]; k++ {
		if _, ok := occupiedCols[k]; !ok {
			colDist += expansionDist
		} else {
			colDist++
		}
	}

	return rowDist + colDist
}

type image []string

func (i image) findGalaxies() []index {
	var galaxies []index
	for i, row := range i {
		for j, r := range row {
			if r == '#' {
				galaxies = append(galaxies, index{i, j})
			}
		}
	}
	return galaxies
}

func (i image) findOccupiedRowsAndCols(galaxies []index) (map[int]struct{}, map[int]struct{}) {
	occupiedRows := make(map[int]struct{})
	occupiedCols := make(map[int]struct{})
	for _, g := range galaxies {
		occupiedRows[g[0]] = struct{}{}
		occupiedCols[g[1]] = struct{}{}
	}
	return occupiedRows, occupiedCols
}
