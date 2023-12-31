package main

import (
	"fmt"

	"github.com/johnliu98/advent-of-code/2023/internal/parse"
	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

func main() {
	histories := read.ReadLines("input.txt")

	var ans int
	for _, h := range histories {
		history := parse.Ints(h)
		ans += extrapolate(history)
	}

	fmt.Println("Answer: ", ans)
}

func extrapolate(values []int) int {
	if allZeros(values) {
		return 0
	}

	diff := make([]int, len(values)-1)
	for i := 0; i < len(values)-1; i++ {
		diff[i] = values[i+1] - values[i]
	}

	diffExt := extrapolate(diff)

	return values[len(values)-1] + diffExt
}

func allZeros(values []int) bool {
	for _, v := range values {
		if v != 0 {
			return false
		}
	}
	return true
}
