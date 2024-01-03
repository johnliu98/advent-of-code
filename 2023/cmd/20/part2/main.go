package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/johnliu98/advent-of-code/2023/internal/read"
)

func main() {
	lines := read.ReadLines("input.txt")

	connections := make(map[string][]string)
	modules := make(map[string]module)
	conjunctions := make([]string, 0)
	for _, line := range lines {
		x := strings.Split(line, " -> ")
		ms := strings.Split(x[1], ", ")
		if x[0] == "broadcaster" {
			modules[x[0]] = newBroadcaster()
			connections[x[0]] = ms
			continue
		}

		name := x[0][1:]
		connections[name] = ms

		if x[0][0] == '%' {
			modules[name] = newFlip()
		} else {
			conjunctions = append(conjunctions, x[0][1:])
		}
	}

	for _, c := range conjunctions {
		var inputs []string
		for mod, conns := range connections {
			for _, conn := range conns {
				if conn == c {
					inputs = append(inputs, mod)
					break
				}
			}
		}
		modules[c] = newConjunction(inputs)
	}

	var final *conjunction
	for tx, conns := range connections {
		if len(conns) == 1 && conns[0] == "rx" {
			final = modules[tx].(*conjunction)
		}
	}

	var inputs []string
	for n := range final.last {
		inputs = append(inputs, n)
	}

	cycles := make([]int, len(inputs))
	for i, done := 1, false; !done; i++ {
		transmitters := []string{"button"}
		receivers := []string{"broadcaster"}
		pulses := []pulse{low}

		for len(transmitters) > 0 && len(receivers) > 0 && len(pulses) > 0 {
			t, r, p := transmitters[0], receivers[0], pulses[0]
			transmitters, receivers, pulses = transmitters[1:], receivers[1:], pulses[1:]

			for k, n := range inputs {
				if final.last[n] == high {
					cycles[k] = i
				}
			}

			m, ok := modules[r]
			if !ok {
				continue
			}

			q := m.send(t, p)
			if q == none {
				continue
			}

			conns := connections[r]

			for _, c := range conns {
				transmitters = append(transmitters, r)
				receivers = append(receivers, c)
				pulses = append(pulses, q)
			}
		}

		done = true
		for _, c := range cycles {
			if c == 0 {
				done = false
			}
		}
	}

	ans := leastCommonMultiple(cycles...)

	fmt.Println("Answer: ", ans)
}

type module interface {
	send(string, pulse) pulse
	state() string
}

type broadcaster struct{}

func newBroadcaster() *broadcaster {
	return &broadcaster{}
}

func (b *broadcaster) send(_ string, p pulse) pulse {
	return p
}

func (b *broadcaster) state() string {
	return ""
}

type flip struct {
	s int
}

func newFlip() *flip {
	return &flip{}
}

func (f *flip) send(_ string, p pulse) pulse {
	if p == high {
		return none
	}
	f.s = f.s ^ 1
	return pulse(f.s)
}

func (f *flip) state() string {
	return strconv.Itoa(f.s)
}

type conjunction struct {
	last map[string]pulse
}

func newConjunction(inputs []string) *conjunction {
	last := make(map[string]pulse)
	for _, in := range inputs {
		last[in] = low
	}
	return &conjunction{last: last}
}

func (c *conjunction) send(n string, p pulse) pulse {
	c.last[n] = p
	for _, l := range c.last {
		if l == low {
			return high
		}
	}
	return low
}

func (c *conjunction) state() string {
	var s int
	for _, l := range c.last {
		if l == high {
			s++
		}
	}
	return strconv.Itoa(s)
}

type pulse int

const (
	none pulse = -1
	low  pulse = 0
	high pulse = 1
)

func (p pulse) String() string {
	switch p {
	case low:
		return "low"
	case high:
		return "high"
	default:
		return "none"
	}
}
