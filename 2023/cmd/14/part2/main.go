package main

import (
	"fmt"
	"strings"

	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

const (
	round = 'O'
	cube  = '#'
	empty = '.'
)

const cycles = 1000000000

func main() {
	platformStrings := read.ReadLines("input.txt")
	platform := make(platform, len(platformStrings))
	for i, rowString := range platformStrings {
		row := make([]rune, len(rowString))
		for j, c := range rowString {
			row[j] = c
		}
		platform[i] = row
	}

	cache := map[string]int{platform.string(): 0}
	var step, cycle int
	for i := 1; i <= cycles; i++ {
		platform.spinCycle()
		if j, ok := cache[platform.string()]; ok {
			step = j
			cycle = i - j
			break
		}
		cache[platform.string()] = i
	}

	var finalPlatformString string
	for p, i := range cache {
		if i == step+(cycles-step)%cycle {
			finalPlatformString = p
			break
		}
	}

	finalPlatform := platform
	for i, rowString := range strings.Split(finalPlatformString, fmt.Sprintln()) {
		row := make([]rune, len(rowString))
		for j, c := range rowString {
			row[j] = c
		}
		platform[i] = row
	}

	ans := finalPlatform.northLoad()

	fmt.Println("Answer: ", ans)
}

type platform [][]rune

func (p platform) string() string {
	var s string
	for _, row := range p {
		for _, c := range row {
			s += string(c)
		}
		s += fmt.Sprintln()
	}
	return s[:len(s)-1]
}

func (p platform) northLoad() int {
	var load int
	for i, row := range p {
		for _, c := range row {
			if c == round {
				load += len(p) - i
			}
		}
	}
	return load
}

func (p platform) spinCycle() {
	p.northTilt()
	p.westTilt()
	p.southTilt()
	p.eastTilt()
}

func (p platform) northTilt() {
	for j := 0; j < len(p[0]); j++ {
		var ri int
		for i := 0; i < len(p); i++ {
			if p[i][j] == round {
				p[i][j] = empty
				p[ri][j] = round
				ri++
			}
			if p[i][j] == cube {
				ri = i + 1
			}
		}
	}
}

func (p platform) westTilt() {
	for i := 0; i < len(p); i++ {
		var rj int
		for j := 0; j < len(p[0]); j++ {
			if p[i][j] == round {
				p[i][j] = empty
				p[i][rj] = round
				rj++
			}
			if p[i][j] == cube {
				rj = j + 1
			}
		}
	}
}

func (p platform) southTilt() {
	for j := 0; j < len(p[0]); j++ {
		ri := len(p) - 1
		for i := len(p) - 1; i >= 0; i-- {
			if p[i][j] == round {
				p[i][j] = empty
				p[ri][j] = round
				ri--
			}
			if p[i][j] == cube {
				ri = i - 1
			}
		}
	}
}

func (p platform) eastTilt() {
	for i := 0; i < len(p); i++ {
		rj := len(p[0]) - 1
		for j := len(p[0]) - 1; j >= 0; j-- {
			if p[i][j] == round {
				p[i][j] = empty
				p[i][rj] = round
				rj--
			}
			if p[i][j] == cube {
				rj = j - 1
			}
		}
	}
}
