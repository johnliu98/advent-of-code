package main

import (
	"fmt"

	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

const (
	round = 'O'
	cube  = '#'
)

func main() {
	platform := read.ReadLines("./cmd/14/input.txt")

	var ans int
	platformLength := len(platform)
	for j := range platform[0] {
		var ri int
		for i := 0; i < platformLength; i++ {
			if platform[i][j] == round {
				ans += platformLength - ri
				ri++
			}
			if platform[i][j] == cube {
				ri = i + 1
			}
		}
	}

	fmt.Println("Answer: ", ans)
}
