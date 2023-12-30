package main

type direction int

const (
	up direction = iota
	left
	down
	right
)

func (d direction) opposite(p direction) bool {
	switch d {
	case up:
		return p == down
	case left:
		return p == right
	case down:
		return p == up
	case right:
		return p == left
	default:
		return false
	}
}
