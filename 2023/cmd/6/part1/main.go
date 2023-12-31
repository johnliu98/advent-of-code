package main

import (
	"fmt"
	"math"

	"github.com/johnliu98/advent-of-code/2023/internal/parse"
	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

func main() {
	lines := read.ReadLines("input.txt")

	times := parse.Ints(lines[0])
	distances := parse.Ints(lines[1])

	ans := 1
	for i := 0; i < len(times); i++ {
		t := float64(times[i])
		d := float64(distances[i])

		dx := math.Sqrt(t*t-4*d) / 2

		minFloat := t/2 - dx
		maxFloat := t/2 + dx

		minRounded := math.Ceil(minFloat)
		maxRounded := math.Floor(maxFloat)

		minInt := int(minRounded)
		if math.Abs(minFloat-minRounded) == 0 {
			minInt++
		}

		maxInt := int(maxRounded)
		if math.Abs(maxFloat-maxRounded) == 0 {
			maxInt--
		}

		wins := maxInt - minInt + 1

		ans *= wins
	}

	fmt.Println("Answer: ", ans)
}
