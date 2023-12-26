package main

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

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
		}

		ans += value
	}

	fmt.Println("Answer: ", ans)
}
