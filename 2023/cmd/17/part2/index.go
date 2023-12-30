package main

type index struct {
	i, j int
}

func (i index) dist(j index) int {
	x := i.i - j.i
	if x < 0 {
		x = -x
	}
	y := i.j - j.j
	if y < 0 {
		y = -y
	}
	return x + y
}

func (i index) move(d direction, count int) index {
	switch d {
	case up:
		return index{i: i.i - count, j: i.j}
	case left:
		return index{i: i.i, j: i.j - count}
	case down:
		return index{i: i.i + count, j: i.j}
	case right:
		return index{i: i.i, j: i.j + count}
	default:
		panic("invalid direction")
	}
}
