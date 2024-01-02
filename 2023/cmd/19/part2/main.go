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

	names := []string{"in"}
	parts := []part{{"x": []int{1, 4000}, "m": []int{1, 4000}, "a": []int{1, 4000}, "s": []int{1, 4000}}}
	acceptedParts := []part{}
	for len(parts) > 0 && len(names) > 0 {
		n, p := names[0], parts[0]
		names, parts = names[1:], parts[1:]

		newNames, newParts := workflows[n].process(p)

		for i := range newParts {
			if newNames[i] == accepted {
				acceptedParts = append(acceptedParts, newParts[i])
			}
			if newNames[i] != rejected {
				parts = append(parts, newParts[i])
				names = append(names, newNames[i])
			}
		}
	}

	var ans int
	for _, p := range acceptedParts {
		x := p["x"][1] - p["x"][0] + 1
		m := p["m"][1] - p["m"][0] + 1
		a := p["a"][1] - p["a"][0] + 1
		s := p["s"][1] - p["s"][0] + 1
		ans += x * m * a * s
	}

	fmt.Println("Answer: ", ans)
}

type workflow []rule

func (w workflow) process(p part) ([]string, []part) {
	var ss []string
	var ps []part
	for _, r := range w {
		if p["x"] == nil {
			break
		}
		if n, a, d := r(p); a["x"] != nil {
			ss = append(ss, n)
			ps = append(ps, a)
			p = d
		}
	}
	return ss, ps
}

type rule func(part) (string, part, part)

func newRule(s string) rule {
	x := strings.Split(s, ":")

	if len(x) < 2 {
		return func(p part) (string, part, part) {
			return x[0], p, part{}
		}
	}

	condition := x[0]

	key := string(condition[0])

	threshold := parse.Int(condition[2:])

	filter := func(p part) (part, part) {
		v := p[key]
		if v[1] < threshold {
			return p, part{}
		}
		if v[0] >= threshold {
			return part{}, p
		}

		p1, p2 := make(part), make(part)
		for k, r := range p {
			if k != key {
				p1[k] = r
				p2[k] = r
				continue
			}

			p1[k] = []int{v[0], threshold - 1}
			p2[k] = []int{threshold, v[1]}
		}

		return p1, p2
	}
	if condition[1] == '>' {
		filter = func(p part) (part, part) {
			v := p[key]
			if v[0] > threshold {
				return p, part{}
			}
			if v[1] <= threshold {
				return part{}, p
			}

			p1, p2 := make(part), make(part)
			for k, r := range p {
				if k != key {
					p1[k] = r
					p2[k] = r
					continue
				}

				p1[k] = []int{threshold + 1, v[1]}
				p2[k] = []int{v[0], threshold}
			}

			return p1, p2
		}
	}

	return func(p part) (string, part, part) {
		a, d := filter(p)
		return x[1], a, d
	}
}

type part map[string][]int
