package main

import (
	"fmt"
	"regexp"

	"github.com/johnliu98/advent-of-code/2023/internal/conv"
	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

func main() {
	lines := read.ReadLines("input.txt")

	var ans int
	// (i, j) of '*' -> part number
	gears := make(map[[2]int][]int)
	re := regexp.MustCompile(`\d+`)
	for i := 0; i < len(lines); i++ {
		partLocs := re.FindAllIndex([]byte(lines[i]), -1)
		for _, partLoc := range partLocs {
			if gearLoc := adjacentStarLocation(lines, i, partLoc); gearLoc[0] != 0 || gearLoc[1] != 0 {
				gearPart := conv.IntFromString(lines[i][partLoc[0]:partLoc[1]])
				gears[gearLoc] = append(gears[gearLoc], gearPart)
			}
		}
	}

	for _, parts := range gears {
		if len(parts) != 2 {
			continue
		}

		ans += parts[0] * parts[1]
	}

	fmt.Println("Answer: ", ans)
}

func adjacentStarLocation(lines []string, row int, loc []int) [2]int {
	for i := row - 1; i <= row+1; i++ {
		if i < 0 || i >= len(lines) {
			continue
		}
		line := lines[i]
		for j := loc[0] - 1; j < loc[1]+1; j++ {
			if j < 0 || j >= len(line) {
				continue
			}
			if line[j] == '*' {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{}
}
