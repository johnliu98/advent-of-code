package main

import (
	"fmt"
	"math"

	"github.com/johnliu98/advent-of-code/2023/internal/conv"
	"github.com/johnliu98/advent-of-code/2023/internal/parse"
	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

func main() {
	lines := read.ReadLines("./cmd/5/input.txt")

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
		category = convertCategory(category, maps)
	}
	location := category

	ans := math.MaxInt64
	for i := 0; i < len(location); i += 2 {
		if loc := location[i]; loc < ans {
			ans = loc
		}
	}

	fmt.Println("Answer: ", ans)
}

func convertCategory(category []int, maps [][]int) []int {
	var convertedCategory []int
	for _, m := range maps {
		var nonConvertedCategory []int
		for i := 0; i < len(category); i += 2 {
			r := category[i : i+2]
			inside, outside := partition(r, m)
			converted := convert(inside, m)

			convertedCategory = append(convertedCategory, converted...)
			nonConvertedCategory = append(nonConvertedCategory, outside...)
		}
		category = nonConvertedCategory
	}
	convertedCategory = append(convertedCategory, category...)
	return convertedCategory
}

func convert(r, m []int) []int {
	if len(r) < 2 {
		return nil
	}
	return []int{
		m[0] - m[1] + r[0],
		r[1],
	}
}

func partition(r, m []int) ([]int, []int) {
	if len(r) != 2 {
		panic(fmt.Sprintf("range is not length 2: range %v", r))
	}

	if r[0] >= m[1] && r[0]+r[1] <= m[1]+m[2] {
		// Category range is entirely contained inside source range.
		return r, nil
	}

	if r[0] < m[1] && r[0]+r[1] > m[1]+m[2] {
		// Category range starts before source range, and ends after source range.
		return []int{
				m[1],
				m[2],
			}, []int{
				r[0],
				m[1] - r[0],
				m[1] + m[2],
				r[0] + r[1] - m[1] - m[2],
			}
	}

	if r[0]+r[1] > m[1] && r[0]+r[1] <= m[1]+m[2] {
		// Category range starts before source range, but ends inside source range.
		return []int{
				m[1],
				r[0] + r[1] - m[1],
			}, []int{
				r[0],
				m[1] - r[0],
			}
	}

	if r[0] >= m[1] && r[0] < m[1]+m[2] {
		// Category range starts inside source range, but ends after soure range.
		return []int{
				r[0],
				m[1] + m[2] - r[0],
			}, []int{
				m[1] + m[2],
				r[0] + r[1] - m[1] - m[2],
			}
	}

	// Category range and source range has not overlap.
	return nil, r
}
