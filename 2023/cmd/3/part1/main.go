package main

import (
	"fmt"
	"regexp"
	"unicode"

	"github.com/johnliu98/advent-of-code/2023/internal/conv"
	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

func main() {
	lines := read.ReadLines("./cmd/3/input.txt")

	var ans int
	re := regexp.MustCompile(`\d+`)
	for i := 0; i < len(lines); i++ {
		locs := re.FindAllIndex([]byte(lines[i]), -1)
		for _, loc := range locs {
			if symbolAdjacent(lines, i, loc) {
				ans += conv.IntFromString(lines[i][loc[0]:loc[1]])
			}
		}
	}

	fmt.Println("Answer: ", ans)
}

func symbolAdjacent(lines []string, row int, loc []int) bool {
	for i := row - 1; i <= row+1; i++ {
		if i < 0 || i >= len(lines) {
			continue
		}
		line := lines[i]
		for j := loc[0] - 1; j < loc[1]+1; j++ {
			if j < 0 || j >= len(line) {
				continue
			}
			if isSymbol(line[j]) {
				return true
			}
		}
	}
	return false
}

func isSymbol(b byte) bool {
	if unicode.IsNumber(rune(b)) {
		return false
	}
	if string(b) == "." {
		return false
	}
	return true
}
