package main

type state int

const numStates = 41

func (s state) move(d direction, count int) state {
	if s.direction() == d {
		return s.step(count)
	}

	return state(d + 1).step(count - 1)
}

func (s state) direction() direction {
	return direction((s - 1) % 4)
}

func (s state) step(count int) state {
	return s + 4*state(count)
}

func (s state) invalid() bool {
	return s < 0 || s >= numStates
}
