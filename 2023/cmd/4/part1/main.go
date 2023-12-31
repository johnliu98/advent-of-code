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
	for _, line := range lines {
		var points int

		numCount := make(map[int]int)
		nums := parse.Ints(strings.Split(line, ":")[1])
		for _, n := range nums {
			numCount[n]++
		}

		for _, count := range numCount {
			if count < 2 {
				continue
			}

			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}

		ans += points
	}

	fmt.Println("Answer: ", ans)
}
