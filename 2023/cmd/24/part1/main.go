package main

import (
	"fmt"
	"math"

	"github.com/johnliu98/advent-of-code/2023/internal/parse"
	"github.com/johnliu98/advent-of-code/2023/internal/read"
	"gonum.org/v1/gonum/spatial/r2"
)

const eps = 1e-6

const (
	minPos = 200000000000000
	maxPos = 400000000000000
)

func main() {
	lines := read.ReadLines("input.txt")
	hailstones := make([]hailstone, len(lines))
	for i, line := range lines {
		ints := parse.Ints(line)
		hailstones[i] = hailstone{
			p: r2.Vec{X: float64(ints[0]), Y: float64(ints[1])},
			v: r2.Vec{X: float64(ints[3]), Y: float64(ints[4])},
		}
	}

	var ans int
	for i := 0; i < len(hailstones); i++ {
		for j := i + 1; j < len(hailstones); j++ {
			p := hailstones[i].intersection(hailstones[j])
			if p.X >= minPos && p.X <= maxPos && p.Y >= minPos && p.Y <= maxPos {
				ans++
			}
		}
	}

	fmt.Println("Answer: ", ans)
}

type hailstone struct {
	p, v r2.Vec
}

// https://en.wikipedia.org/wiki/Line%E2%80%93line_intersection#Given_two_points_on_each_line_segment
func (h hailstone) intersection(s hailstone) r2.Vec {
	p1 := h.p
	p2 := r2.Add(h.p, h.v)
	p3 := s.p
	p4 := r2.Add(s.p, s.v)

	den := (p1.X-p2.X)*(p3.Y-p4.Y) - (p1.Y-p2.Y)*(p3.X-p4.X)
	if math.Abs(den) < eps {
		return r2.Vec{}
	}

	t := (p1.X-p3.X)*(p3.Y-p4.Y) - (p1.Y-p3.Y)*(p3.X-p4.X)
	t /= den
	if t < 0 {
		return r2.Vec{}
	}

	u := (p1.X-p3.X)*(p1.Y-p2.Y) - (p1.Y-p3.Y)*(p1.X-p2.X)
	u /= den
	if u < 0 {
		return r2.Vec{}
	}

	x := p1.X + t*(p2.X-p1.X)
	y := p1.Y + t*(p2.Y-p1.Y)

	return r2.Vec{X: x, Y: y}
}
