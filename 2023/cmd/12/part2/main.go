package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/johnliu98/advent-of-code/2023/internal/parse"
	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

const (
	operational = '.'
	damaged     = '#'
	unknown     = '?'
)

const numFolds = 5

var cache = make(map[string]int)

func hash(springs string, groups []int) string {
	h := springs
	for _, g := range groups {
		h += strconv.Itoa(g)
	}
	return h
}

func main() {
	records := read.ReadLines("input.txt")

	var ans int
	for _, record := range records {
		split := strings.Split(record, " ")

		springs := split[0]
		groups := parse.Ints(split[1])
		for i := 0; i < numFolds-1; i++ {
			springs += string(unknown) + split[0]
			groups = append(groups, parse.Ints(split[1])...)
		}

		ans += numValid(springs, groups)
	}

	fmt.Println("Answer: ", ans)
}

func numValid(springs string, groups []int) int {
	if n, ok := cache[hash(springs, groups)]; ok {
		return n
	}

	valids := make([]int, 0)
	g := groups[0]

	m := len(springs) - len(groups) + 1
	for _, g := range groups {
		m -= g
	}

	for si := 0; si <= m; si++ {
		if si-1 >= 0 && springs[si-1] == damaged {
			break
		}
		if si+g < len(springs) && springs[si+g] == damaged {
			continue
		}

		if len(groups) == 1 {
			var damagedCount int
			for _, s := range springs[si+g:] {
				if s == damaged {
					damagedCount++
				}
			}
			if damagedCount > 0 {
				continue
			}
		}

		ok := true
		for sj := si; sj < si+g; sj++ {
			if springs[sj] == operational {
				ok = false
				break
			}
		}
		if ok {
			valids = append(valids, si)
		}
	}

	if len(groups) == 1 {
		return len(valids)
	}

	var count int
	for _, c := range valids {
		siNext := c + g + 1
		for si := siNext; si < len(springs); si++ {
			if springs[si] != operational {
				siNext = si
				break
			}
		}
		count += numValid(springs[siNext:], groups[1:])
	}

	cache[hash(springs, groups)] = count

	return count
}
