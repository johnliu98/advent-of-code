package main

import (
	"fmt"
	"strings"

	"github.com/johnliu98/advent-of-code/2023/internal/parse"
	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

func main() {
	lines := read.ReadLines("input.txt")

	var ans int
	copies := make([]int, len(lines))
	for i, line := range lines {
		var matches int

		numCount := make(map[int]int)
		nums := parse.Ints(strings.Split(line, ":")[1])
		for _, n := range nums {
			numCount[n]++
		}

		for _, count := range numCount {
			if count < 2 {
				continue
			}
			matches++
		}

		ans += 1 + copies[i]

		for j := 0; j < matches; j++ {
			copies[i+j+1] += 1 + copies[i]
		}
	}

	fmt.Println("Answer: ", ans)
}
