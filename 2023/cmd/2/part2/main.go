package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/johnliu98/advent-of-code/2023/internal/parse"
	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

func main() {
	input := read.ReadLines("input.txt")

	var ans int
	for _, in := range input {
		colors := map[string]int{}
		for _, s := range parse.SplitValues(in, ";") {
			for _, cube := range strings.Split(s, ",") {
				x := strings.Split(strings.Trim(cube, " "), " ")

				n, err := strconv.Atoi(x[0])
				if err != nil {
					panic(err)
				}
				color := x[1]

				colors[color] = max(colors[color], n)
			}
		}

		ans += colors["red"] * colors["green"] * colors["blue"]
	}

	fmt.Println("Answer: ", ans)
}
