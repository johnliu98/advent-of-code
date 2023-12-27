package main

import (
	"fmt"
	"strings"

	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

func main() {
	initializationSequence := read.ReadLines("input.txt")[0]
	steps := strings.Split(initializationSequence, ",")

	var ans int
	for _, s := range steps {
		ans += hash(s)
	}

	fmt.Println("Answer: ", ans)
}

func hash(s string) int {
	var h uint8
	for _, c := range s {
		h += uint8(c)
		h *= 17
	}
	return int(h)
}
