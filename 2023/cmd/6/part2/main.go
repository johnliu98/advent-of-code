package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

func main() {
	lines := read.ReadLines("input.txt")

	t := parseNumber(lines[0])
	d := parseNumber(lines[1])

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

	fmt.Println("Answer: ", wins)
}

func parseNumber(s string) float64 {
	re := regexp.MustCompile(`\d+`)
	subNums := re.FindAll([]byte(s), -1)
	var numString string
	for _, n := range subNums {
		numString += string(n)
	}

	i, err := strconv.Atoi(numString)
	if err != nil {
		panic(err)
	}

	return float64(i)
}
