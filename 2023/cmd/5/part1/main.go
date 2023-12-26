package main

import (
	"fmt"
	"slices"

	"github.com/johnliu98/advent-of-code/2023/internal/conv"
	"github.com/johnliu98/advent-of-code/2023/internal/parse"
	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

func main() {
	lines := read.ReadLines("input.txt")

	seeds := conv.IntsFromString(lines[0])

	blocks := parse.Blocks(lines[1:])
	mapsPerCategory := make([][][]int, len(blocks))
	for i, block := range blocks {
		for _, m := range block[1:] {
			mapsPerCategory[i] = append(mapsPerCategory[i], conv.IntsFromString(m))
		}
	}

	category := seeds
	for _, maps := range mapsPerCategory {
		convertCategory(category, maps)
	}
	location := category

	fmt.Println("Answer: ", slices.Min(location))
}

func convertCategory(category []int, maps [][]int) {
	for i := range category {
		var converted bool
		for _, m := range maps {
			if converted {
				break
			}
			category[i], converted = convert(category[i], m)
		}
	}
}

func convert(n int, m []int) (int, bool) {
	d := n - m[1]
	if d < 0 || d >= m[2] {
		return n, false
	}
	return m[0] + d, true
}
