package main

type index struct {
	i, j int
}

func (i index) move(d direction) index {
	switch d {
	case up:
		return index{i: i.i - 1, j: i.j}
	case left:
		return index{i: i.i, j: i.j - 1}
	case down:
		return index{i: i.i + 1, j: i.j}
	case right:
		return index{i: i.i, j: i.j + 1}
	default:
		panic("invalid direction")
	}
}
