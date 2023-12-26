package main

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

var numbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	input := read.ReadLines("input.txt")

	var ans int
	for _, in := range input {
		var value int

		for i := 0; i < len(in); i++ {
			char := rune(in[i])
			if unicode.IsNumber(char) {
				n, err := strconv.Atoi(string(char))
				if err != nil {
					panic(err)
				}
				value += 10 * n
				break
			}
			if n := getNumber(in[i:]); n != 0 {
				value += 10 * n
				break
			}
		}

		for i := len(in) - 1; i >= 0; i-- {
			char := rune(in[i])
			if unicode.IsNumber(char) {
				n, err := strconv.Atoi(string(char))
				if err != nil {
					panic(err)
				}
				value += n
				break
			}
			if n := getNumber(in[i:]); n != 0 {
				value += n
				break
			}
		}

		ans += value
	}

	fmt.Println("Answer: ", ans)
}

func getNumber(s string) int {
	if len(s) < 3 {
		return 0
	}

	if n, ok := numbers[s[:3]]; ok {
		return n
	}

	if len(s) < 4 {
		return 0
	}

	if n, ok := numbers[s[:4]]; ok {
		return n
	}

	if len(s) < 5 {
		return 0
	}

	return numbers[s[:5]]
}
