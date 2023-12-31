package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/johnliu98/advent-of-code/2023/internal/parse"
	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

func main() {
	input := read.ReadLines("input.txt")

	var ans int
	for _, in := range input {
		valid := true

		sets := strings.Split(strings.Split(in, ":")[1], ";")
		for _, s := range sets {
			colors := map[string]int{}
			cubes := strings.Split(s, ",")
			for _, c := range cubes {
				x := strings.Split(strings.Trim(c, " "), " ")
				n, err := strconv.Atoi(x[0])
				if err != nil {
					panic(err)
				}
				colors[x[1]] = n
			}

			if colors["red"] > maxRed || colors["green"] > maxGreen || colors["blue"] > maxBlue {
				valid = false
			}
		}

		if valid {
			ans += parse.Int(strings.Split(strings.Split(in, ":")[0], " ")[1])
		}
	}

	fmt.Println("Answer: ", ans)
}
