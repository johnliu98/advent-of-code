package main

type state int

const numStates = 13

const (
	none state = iota
	oneUp
	oneLeft
	oneDown
	oneRight
	twoUp
	twoLeft
	twoDown
	twoRight
	threeUp
	threeLeft
	threeDown
	threeRight
)

func (s state) move(d direction) state {
	if s.direction() == d {
		return s.step()
	}

	return state(d + 1)
}

func (s state) direction() direction {
	return direction((s - 1) % 4)
}

func (s state) step() state {
	return s + 4
}

func (s state) invalid() bool {
	return s < 0 || s >= numStates
}
