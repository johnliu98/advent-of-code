package main

import (
	"fmt"

	"github.com/johnliu98/advent-of-code/2023/internal/parse"
	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

func main() {
	lines := read.ReadLines("./cmd/13/input.txt")
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
	for i, j := row-1, row; i >= 0 && j < len(pattern); i, j = i-1, j+1 {
		if pattern[i] != pattern[j] {
			return false
		}
	}
	return true
}

func isVerticalReflection(pattern []string, col int) bool {
	for i, j := col-1, col; i >= 0 && j < len(pattern[0]); i, j = i-1, j+1 {
		for _, row := range pattern {
			if row[i] != row[j] {
				return false
			}
		}
	}
	return true
}
