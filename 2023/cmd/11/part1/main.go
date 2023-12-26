package main

import (
	"fmt"

	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

func main() {
	image := image(read.ReadLines("input.txt"))

	galaxies := image.findGalaxies()
	occupiedRows, occupiedCols := image.findOccupiedRowsAndCols(galaxies)

	expandedImage := expandImage(image, occupiedRows, occupiedCols)
	expandedGalaxies := expandedImage.findGalaxies()

	var ans int
	for i := 0; i < len(expandedGalaxies)-1; i++ {
		for j := i + 1; j < len(expandedGalaxies); j++ {
			ans += expandedGalaxies[i].dist(expandedGalaxies[j])
		}
	}

	fmt.Println("Answer: ", ans)
}

func expandImage(originalImage image, occupiedRows, occupiedCols map[int]struct{}) image {
	var expandedImage image
	for i, row := range originalImage {
		var expandedRow string
		for j, r := range row {
			expandedRow += string(r)
			if _, ok := occupiedCols[j]; !ok {
				expandedRow += string(r)
			}
		}
		expandedImage = append(expandedImage, expandedRow)
		if _, ok := occupiedRows[i]; !ok {
			expandedImage = append(expandedImage, expandedRow)
		}
	}
	return expandedImage
}

type index [2]int

func (i index) dist(j index) int {
	x := j[1] - i[1]
	if x < 0 {
		x = -x
	}
	y := j[0] - i[0]
	if y < 0 {
		y = -y
	}
	return x + y
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
