package main

import (
	"fmt"
	"strings"

	"github.com/johnliu98/advent-of-code/2023/internal/parse"
	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

const (
	accepted = "A"
	rejected = "R"
)

func main() {
	lines := read.ReadLines("input.txt")
	blocks := parse.Blocks(lines)

	workflows := make(map[string]workflow)
	for _, s := range blocks[0] {
		x := strings.Split(strings.TrimRight(s, "}"), "{")
		name := x[0]
		var w workflow
		for _, s := range strings.Split(x[1], ",") {
			w = append(w, newRule(s))
		}

		workflows[name] = w
	}

	var ans int
	for _, s := range blocks[1] {
		x := parse.Ints(s)
		p := part{x: x[0], m: x[1], a: x[2], s: x[3]}

		w := "in"
		for w != accepted && w != rejected {
			w = workflows[w].process(p)
		}

		if w == accepted {
			ans += p.x + p.m + p.a + p.s
		}
	}

	fmt.Println("Answer: ", ans)
}

type workflow []rule

func (w workflow) process(p part) string {
	for _, rule := range w {
		if res := rule(p); res != "" {
			return res
		}
	}
	return ""
}

type rule func(part) string

func newRule(s string) rule {
	x := strings.Split(s, ":")

	if len(x) < 2 {
		return func(part) string { return x[0] }
	}

	condition := x[0]

	value := func(p part) int {
		switch condition[0] {
		case 'x':
			return p.x
		case 'm':
			return p.m
		case 'a':
			return p.a
		case 's':
			return p.s
		default:
			panic("invalid letter")
		}
	}

	threshold := parse.Int(condition[2:])

	compare := func(x int) bool {
		return x < threshold
	}
	if condition[1] == '>' {
		compare = func(x int) bool {
			return x > threshold
		}
	}

	return func(p part) string {
		if compare(value(p)) {
			return x[1]
		}
		return ""
	}
}

type part struct {
	x, m, a, s int
}
