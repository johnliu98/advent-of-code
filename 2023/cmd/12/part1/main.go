package main

import (
	"fmt"
	"strings"

	"github.com/johnliu98/advent-of-code/2023/internal/parse"
	"github.com/johnliu98/advent-of-code/2023/internal/read"
	"gonum.org/v1/gonum/stat/combin"
)

const (
	operational = '.'
	damaged     = '#'
	unknown     = '?'
)

func main() {
	records := read.ReadLines("./cmd/12/input.txt")

	var ans int
	for _, record := range records {
		split := strings.Split(record, " ")
		springs := split[0]
		groups := parse.Numbers(split[1])

		numUnknownDamaged := sum(groups)
		var unknowns []int
		for i, s := range springs {
			if s == unknown {
				unknowns = append(unknowns, i)
			}
			if s == damaged {
				numUnknownDamaged--
			}
		}

		var numValid int
		combs := combin.Combinations(len(unknowns), numUnknownDamaged)
		for _, c := range combs {
			candidate := candidateFromComb(springs, unknowns, c)
			if valid(candidate, groups) {
				numValid++
			}
		}

		ans += numValid
	}

	fmt.Println("Answer: ", ans)
}

func valid(candidate string, groups []int) bool {
	candidatesPadded := string(operational) + candidate + string(operational)
	var start int
	var candidateGroups []int
	prevSpring := operational
	for i, s := range candidatesPadded {
		if s == damaged && prevSpring == operational {
			start = i
		}
		if s == operational && prevSpring == damaged {
			candidateGroups = append(candidateGroups, i-start)
		}
		prevSpring = s
	}

	if len(candidateGroups) != len(groups) {
		return false
	}

	for i := range candidateGroups {
		if candidateGroups[i] != groups[i] {
			return false
		}
	}

	return true
}

func candidateFromComb(springs string, unknowns []int, comb []int) string {
	springIsDamaged := make([]bool, len(springs))
	for _, i := range comb {
		springIsDamaged[unknowns[i]] = true
	}

	var candidate string
	for i, s := range springs {
		if s == unknown {
			if springIsDamaged[i] {
				candidate += string(damaged)
			} else {
				candidate += string(operational)
			}
		} else {
			candidate += string(s)
		}
	}

	return candidate
}

func sum(values []int) int {
	var sum int
	for _, v := range values {
		sum += v
	}
	return sum
}
