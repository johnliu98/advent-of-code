package main

import (
	"fmt"

	"github.com/johnliu98/advent-of-code/2023/internal/parse"
	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

func main() {
	lines := read.ReadLines("input.txt")
	patterns := parse.Blocks(lines)

	var ans int
	for _, pattern := range patterns {
		for row := 1; row < len(pattern); row++ {
			if isHorizontalReflection(pattern, row) {
				ans += 100 * row
			}
		}
		for col := 1; col < len(pattern[0]); col++ {
			if isVerticalReflection(pattern, col) {
				ans += col
			}
		}
	}

	fmt.Println("Answer: ", ans)
}

func isHorizontalReflection(pattern []string, row int) bool {
	var smudges int
	for i, j := row-1, row; i >= 0 && j < len(pattern); i, j = i-1, j+1 {
		for col := range pattern[0] {
			if pattern[i][col] != pattern[j][col] {
				smudges++
			}
			if smudges > 1 {
				return false
			}
		}
	}

	return smudges == 1
}

func isVerticalReflection(pattern []string, col int) bool {
	var smudges int
	for i, j := col-1, col; i >= 0 && j < len(pattern[0]); i, j = i-1, j+1 {
		for _, row := range pattern {
			if row[i] != row[j] {
				smudges++
			}
			if smudges > 1 {
				return false
			}
		}
	}

	return smudges == 1
}
